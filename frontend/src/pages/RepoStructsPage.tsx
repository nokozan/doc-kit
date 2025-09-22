import React, { useEffect, useState } from "react";
import axios from "axios";
import { useParams } from "react-router-dom";

interface Field {
  name: string;
  type: string;
  tag?: string;
  comment?: string;
}

interface Struct {
  id: number;
  name: string;
  comment?: string;
  fields: Field[];
}

const RepoStructsPage: React.FC = () => {
  const { repoId } = useParams<{ repoId: string }>();
  const [structs, setStructs] = useState<Struct[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    if (!repoId) return;
    setLoading(true);
    axios
      .get<Struct[]>(`/api/repos/${repoId}/structs`)
      .then((response) => {
        setStructs(response.data);
        setLoading(false);
      })
      .catch((error) => {
        console.error("Error fetching structs:", error);
        setLoading(false);
      });
  }, [repoId]);

  return (
    <div className="p-6 space-y-4">
      <h1 className="text-2xl font-bold mb-4">
        Structs in Repository {repoId}
      </h1>
      {loading && <p className="text-gray-600">Loading...</p>}
      {structs.map((struct) => (
        <div
          key={struct.id}
          className="border rounded-lg p-4 shadow-sm hover:shadow-md transition-shadow"
        >
          <h2 className="text-xl font-semibold mb-2">{struct.name}</h2>
          {struct.comment && (
            <p className="text-sm text-gray-600 mb-2">{struct.comment}</p>
          )}
          <ul className="text-sm text-gray-800 mt-2 list-disc list-inside pl-4">
            {struct.fields.map((field, index) => (
              <li key={index} className="mb-1">
                <strong>{field.name}</strong>: {field.type}
                {field.tag && (
                  <span className="text-gray-500"> [{field.tag}]</span>
                )}
                {field.comment && (
                  <p className="text-gray-600 ml-4">{field.comment}</p>
                )}
              </li>
            ))}
          </ul>
        </div>
      ))}
    </div>
  );
};

export default RepoStructsPage;
