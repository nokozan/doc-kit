import React, { useState } from "react";
import axios from "axios";
import { toast } from "sonner";
import { Link } from "react-router-dom";

interface RepoCardProps {
  id: number;
  alias: string;
  url: string;
  branch: string;
  description?: string;
  onSynced?: () => void; //optional callback when sync is done to trigger refresh
}

const RepoCard: React.FC<RepoCardProps> = ({
  id,
  alias,
  url,
  branch,
  description,
  onSynced,
}) => {
  const [loading, setLoading] = useState(false);

  const handleSync = async () => {
    setLoading(true);
    try {
      await axios.post(`/api/repos/${id}/sync`);
      toast.success("Repository synced successfully");
      onSynced?.();
    } catch (error) {
      console.error("Error syncing repository:", error);
      toast.error("Failed to sync repository");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="border rounded-lg p-4 shadow-sm hover:shadow-md transition-shadow">
      <h2 className="text-xl font-semibold mb-2">{alias}</h2>
      <button
        onClick={handleSync}
        disabled={loading}
        className={`px-4 py-2 rounded ${
          loading
            ? "bg-gray-400 cursor-not-allowed"
            : "bg-blue-500 hover:bg-blue-600"
        } text-white`}
      >
        {loading ? "Syncing..." : "Sync"}
      </button>
      <p className="text-sm text-gray-600 mt-2">
        <strong>URL:</strong> {url}
      </p>
      <p className="text-sm text-gray-600">
        <strong>Branch:</strong> {branch}
      </p>
      {description && (
        <p className="text-sm text-gray-600 mt-2">{description}</p>
      )}
      <Link
        to={`/repo/${id}/structs`}
        className="inline-block mt-4 text-blue-500 hover:underline"
      >
        View Structs
      </Link>
    </div>
  );
};

export default RepoCard;
