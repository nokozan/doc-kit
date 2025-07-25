import {Link } from 'react-router-dom';

const dummyStructs = [
  {
    struct: "Struct1",
    methods: ["Validate", "GetInfo", "ToJSON"]
  },
  {
    struct: "Struct2",
    methods: ["CalculateFee", "GetSummary"]
  },
  {
    struct: "Struct3",
    methods: ["GetPrice", "Serialize"]
  }
]


export default function Sidebar() {
    return (
        <aside className="w-64 h-screen bg-slate-100 p-4 border-r overflow-y-auto">
            <h2 className="text-lg font-semibold mb-4">Structs</h2>
            <ul className="space-y-2">
                {dummyStructs.map((item, index) => (
                    <li key={index} className="hover:bg-gray-700 p-2 rounded">
                        <Link to={`/struct/${item.struct.toLowerCase()}`} className="block">
                            {item.struct}
                        </Link>
                        <ul className="ml-4 mt-1 text-gray-600 space-y-1">
                            {item.methods.map((method, methodIndex) => (
                                <li key={methodIndex} className="text-sm text-gray-300">
                                    <Link to={`/struct/${item.struct.toLowerCase()}/${method.toLowerCase()}/methods/${method.toLowerCase()}`} 
                                    className="hover:text-white">
                                        {method}()
                                    </Link>
                                </li>
                            ))}
                        </ul>
                    </li>
                ))}
            </ul>
        </aside>
    );
}