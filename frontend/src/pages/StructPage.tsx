import { Link, useParams } from "react-router-dom";

const dummyStructDetail = {
  name: "User",
  doc: "This struct represents a user in the system.",
  fields: [
    {
      name: "ID",
      type: "int",
      tag: '`json:"id"`',
      comment: "Unique identifier",
    },
    {
      name: "Email",
      type: "string",
      tag: '`json:"email"`',
      comment: "Email address",
    },
    {
      name: "CreatedAt",
      type: "time.Time",
      tag: '`json:"created_at"`',
      comment: "Creation time",
    },
  ],
  methods: ["validate", "GetInfo", "ToJSON"],
};

export default function StructPage() {
  const { name } = useParams();

  return (
    <div className="p-6 space-y-6">
      <h1 className="text-3xl font-bold mb-4">{dummyStructDetail.name} </h1>
      <p className="text-gray-600">{dummyStructDetail.doc}</p>

      <div>
        <h2 className="text-xl font-semibold mb-2">Fields</h2>
        <table className="min-w-full bg-white shadow-md rounded">
          <thead className="bg-gray-200">
            <tr>
              <th className="border px-2 py-2">Field Name</th>
              <th className="border px-2 py-2">Type</th>
              <th className="border px-2 py-2">Tag</th>
              <th className="border px-2 py-2">Comment</th>
            </tr>
          </thead>
          <tbody>
            {dummyStructDetail.fields.map((field) => (
              <tr key={field.name}>
                <td className="border px-2 py-2">{field.name}</td>
                <td className="border px-2 py-2">{field.type}</td>
                <td className="border px-2 py-2 font-mono text-xs text-blue-600">
                  {field.tag}
                </td>
                <td className="border px-2 py-2">{field.comment}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <div>
        <h2 className="text-xl font-semibold mb-2">Methods</h2>
        <ul className="list-disc pl-5 space-y-1">
          {dummyStructDetail.methods.map((method) => (
            <li key={method}>
              <Link
                to={`/struct/${name}/method/${method.toLowerCase}`}
                className="text-blue-600 hover:underline"
              >
                {method}()
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
