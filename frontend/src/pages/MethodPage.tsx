import { useParams } from "react-router-dom";
import { useState } from "react";
import MonacoEditor from "@monaco-editor/react";

export default function MethodPage() {
  const { name, method } = useParams();
  const [input, setInput] = useState(
    '{\n "id": "123",\n "email": "test@example.com"\n}'
  );
  const [output, setOutput] = useState("");

  //Dummy mock execution function
  function runMethod() {
    if (method === "validate") {
      if (input.includes("@")) {
        setOutput("Validation successful!");
      } else {
        setOutput("Validation failed: Invalid email format.");
      }
    } else {
      setOutput(`Output of ${method} for struct ${name}: ${input}`);
    }
  }

  return (
    <div className="p-6 space-y-4">
      <h1 className="text-2xl font-bold mb-4">Method: {method}()</h1>
      <p className="text-gray-600">
        Attached to struct: <strong>{name}</strong>
      </p>

      <div>
        <h2 className="font-semibold mb-2">Input JSON</h2>
        <MonacoEditor
          language="json"
          value={input}
          height="200px"
          theme-="vs-dark"
          options={{ fontSize: 14, minimap: { enabled: false } }}
          onChange={(v) => setInput(v ?? "")}
        />
      </div>
      <button
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        onClick={runMethod}
      >
        Run Method
      </button>

      <div className="mt-4 bg-slate-100 p-4 rounded">
        <h3 className="font-semibold mb-1">Output</h3>
        <pre className="text-sm text-gray-800">{output}</pre>
      </div>
    </div>
  );
}
