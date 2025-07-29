import { useParams } from "react-router-dom";
import { useState } from "react";
import MonacoEditor from "@monaco-editor/react";

// import TestCaseManager from "@/components/TestCaseManager";
import axios from "axios";

import { extractErrorMessage } from "@/utils/http";
import ReactJson from "react-json-view";

export default function MethodPage() {
  const { structName, methodName } = useParams();
  const [input, setInput] = useState("{}");
  const [result, setResult] = useState<unknown>(null);
  const [error, setError] = useState<string | null>(null);

  // const [output, setOutput] = useState("");

  const runMethod = async () => {
    try {
      const inputJson = JSON.parse(input);
      const res = await axios.post<{ result: unknown }>(
        `/run/${structName}/${methodName}`,
        inputJson
      );
      setResult(res.data.result);
      setError(null);
    } catch (err: unknown) {
      setError(extractErrorMessage(err));
      setResult(null);
    }
  };

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">
        {structName} → {methodName}
      </h1>

      <div className="mb-4">
        <label className="text-sm text-gray-500">Input JSON</label>
        <MonacoEditor
          height="200px"
          defaultLanguage="json"
          value={input}
          onChange={(v) => setInput(v || "")}
          theme="vs-dark"
        />
      </div>

      <button
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        onClick={runMethod}
      >
        Try It Out
      </button>

      <div className="mt-6">
        {error && (
          <div className="text-red-500">
            <strong>Error:</strong> {error}
          </div>
        )}
        {result !== null && result !== undefined && (
          <div>
            <h2 className="font-bold text-lg mb-2">Result</h2>
            <ReactJson
              // src={result as object}
              src={result}
              name={false}
              collapsed={false}
              enableClipboard={false}
            />
          </div>
        )}
      </div>
    </div>
  );
}
