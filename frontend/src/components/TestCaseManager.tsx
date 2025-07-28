import { useState, useEffect } from "react";
import MonacoEditor from "@monaco-editor/react";

type TestCase = {
  id: string;
  input: string;
  expected: string;
  actual?: string;
  passed?: boolean;
};

type Props = {
  structName: string;
  methodName: string;
  runMethod: (input: string) => Promise<string>;
};

export default function TestCaseManager({
  structName,
  methodName,
  runMethod,
}: Props) {
  const storageKey = `$tests:${structName}.${methodName}`;
  const [tests, setTests] = useState<TestCase[]>([]);

  useEffect(() => {
    const saved = localStorage.getItem(storageKey);
    if (saved) {
      setTests(JSON.parse(saved));
    }
  }, [storageKey]);

  const saveToStorage = (next: TestCase[]) => {
    localStorage.setItem(storageKey, JSON.stringify(next));
    setTests(next);
  };

  const addNew = () => {
    const newTest: TestCase = {
      id: crypto.randomUUID(),
      input: `{\n "example": "value"\n}`,
      expected: "Expected output here",
    };
    saveToStorage([...tests, newTest]);
  };

  const updateTest = (
    id: string,
    field: "input" | "expected",
    value: string
  ) => {
    const updated = tests.map((t) =>
      t.id === id ? { ...t, [field]: value } : t
    );
    saveToStorage(updated);
  };

  const runOne = async (id: string) => {
    const found = tests.find((t) => t.id === id);
    if (!found) return;
    const actual = await runMethod(found.input);
    const passed = actual.trim() === found.expected.trim();
    const updated = tests.map((t) =>
      t.id === id ? { ...t, actual, passed } : t
    );
    saveToStorage(updated);
  };

  return (
    <div className="space-y-6 mt-10">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-bold">Unit Tests</h2>
        <button
          onClick={addNew}
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          Add Test Case
        </button>
      </div>

      {tests.map((test) => (
        <div key={test.id} className="bg-white p-4 rounded shadow">
          <div className="flex justify-between items-center">
            <span className="text-sm font-medium text-gray-600">
              Test ID: {test.id}
            </span>
            {test.passed !== undefined && (
              <span
                className={`text-sm font-semibold ${
                  test.passed ? "text-green-600" : "text-red-600"
                }`}
              >
                {test.passed ? "Passed" : "Failed"}
              </span>
            )}
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Input JSON
            </label>
            <MonacoEditor
              language="json"
              value={test.input}
              height="100px"
              theme="vs-dark"
              options={{ fontSize: 14, minimap: { enabled: false } }}
              onChange={(v) => updateTest(test.id, "input", v ?? "")}
            />
          </div>
          <div>
            <label className="font-medium text-gray-700 mb-1">
              Expected Output
            </label>
            <MonacoEditor
              language="json"
              value={test.expected}
              height="100px"
              theme="vs-dark"
              options={{ fontSize: 14, minimap: { enabled: false } }}
              onChange={(v) => updateTest(test.id, "expected", v ?? "")}
            />
          </div>
          <button
            className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 mt-2"
            onClick={() => runOne(test.id)}
          >
            Run Test
          </button>
          {test.actual !== undefined && (
            <div className="bg-slate-100 p-4 rounded mt-2">
              <strong>Actual:</strong>
              <pre>{test.actual}</pre>
            </div>
          )}
        </div>
      ))}
    </div>
  );
}
