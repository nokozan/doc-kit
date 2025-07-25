import { useParams } from "react-router-dom";


export default function StructPage() {
    const { name } = useParams();
    
    return (
        <div className="p-4">
        <h1 className="text-2xl font-bold mb-4">Struct </h1>
        <p className="text-lg">You are viewing the struct: {name}</p>
        </div>
    );
}
