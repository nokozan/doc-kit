// import { useState } from "react";
// import reactLogo from "./assets/react.svg";
// import viteLogo from "/vite.svg";
import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Sidebar from "./components/Sidebar";
import StructPage from "./pages/StructPage";
import MethodPage from "./pages/MethodPage";
import Home from "./pages/Home";
import { Toaster } from "sonner";
import RepoListPage from "./pages/RepoListPage";
import RepoStructsPage from "./pages/RepoStructsPage";

function App() {
  // const [count, setCount] = useState(0)

  return (
    <Router>
      <Toaster position="top-right" richColors />
      <div className="flex h-screen">
        <div className="w-64 bg-slate-100 border-r p-4">
          <Sidebar />
        </div>

        <div className="flex-1 p-6 overflow-y-auto">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/repos" element={<RepoListPage />} />
            <Route path="/repo/:repoId/structs" element={<RepoStructsPage />} />
            <Route path="/struct/:name" element={<StructPage />} />
            <Route
              path="/struct/:name/method/:method"
              element={<MethodPage />}
            />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
