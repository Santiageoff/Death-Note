import React, { Suspense } from "react";
import { fetchData } from "./hooks/fetchData";
import "./App.css";

const apiData = fetchData("http://localhost:8080/personas");

function App() {
  const data = apiData.read();

  return (
    <div className="App">
      <h1>Death Note </h1>
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
