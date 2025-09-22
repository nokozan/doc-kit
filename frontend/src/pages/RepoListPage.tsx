import React, { useEffect, useState } from "react";
import axios from "axios";
import RepoCard from "@/components/RepoCard";

interface Repo {
  id: number;
  alias: string;
  url: string;
  branch: string;
  description?: string;
}

const RepoListPage: React.FC = () => {
  const [repos, setRepos] = useState<Repo[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  const fetchRepos = async () => {
    setLoading(true);
    try {
      const response = await axios.get<Repo[]>("/api/repos");
      setRepos(response.data);
    } catch (error) {
      console.error("Error fetching repositories:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchRepos();
  }, []);

  return (
    <div className="p-6 space-y-4">
      <h1 className="text-2xl font-bold mb-4">Repositories</h1>
      {loading && <p className="text-gray-600">Loading...</p>}
      {repos.map((repo) => (
        <RepoCard
          key={repo.id}
          {...repo}
          onSynced={fetchRepos} // Refresh list after sync
        />
      ))}
    </div>
  );
};

export default RepoListPage;
