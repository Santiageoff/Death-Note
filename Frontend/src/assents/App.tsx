import React, { Suspense } from "react";
import { fetchData } from "./hooks/fetchData";
import "./App.css";

// Cambiá esto por tu endpoint real en Go
const apiData = fetchData("http://localhost:8080/personas");

function App() {
  const data = apiData.read();

  return (
    <div className="App">
      <h1>Usuarios desde Go</h1>
      <Suspense fallback={<div>Loading...</div>}>
        <ul className="card">
          {data?.map((user: any) => (
            <li key={user.id}>{user.name}</li>
          ))}
        </ul>
      </Suspense>
    </div>
  );
}

export default App;
