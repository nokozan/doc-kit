// import { useState } from "react";
// import reactLogo from "./assets/react.svg";
// import viteLogo from "/vite.svg";
import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Sidebar from "./components/Sidebar";
import StructPage from "./pages/StructPage";
import MethodPage from "./pages/MethodPage";
import Home from "./pages/Home";

function App() {
  // const [count, setCount] = useState(0)

  return (
    <Router>
      <div className="flex h-screen">
        <div className="w-64 bg-slate-100 border-r p-4">
          <Sidebar />
        </div>

        <div className="flex-1 p-6 overflow-y-auto">
          <Routes>
            <Route path="/" element={<Home />} />
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
