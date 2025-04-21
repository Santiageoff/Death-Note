import React, { Suspense, useState } from "react";
import { fetchData } from "./hooks/fetchData";
import FormularioMuerte from "./components/pages/formulario";
import "./App.css";

const apiData = fetchData("http://localhost:8080/personas");

function VerMuertes() {
  const data = apiData.read();

  return (
    <div className="muertes">
      <h1>Lista de Muertes</h1>
      <div className="lista">
        {data?.map((user: any) => (
          <div key={user.id} className="persona">
            <img src={user.foto_url} alt={`${user.nombre} ${user.apellido}`} width={150} />
            <h2>{user.nombre} {user.apellido}</h2>
            <p><strong>Causa:</strong> {user.causa_muerte}</p>
            <p><strong>Fecha:</strong> {user.fecha_muerte}</p>
          </div>
        ))}
      </div>
    </div>
  );
}


function App() {
  const [vista, setVista] = useState<"formulario" | "muertes">("formulario");

  return (
    <div className="App">
      <h1>Death Note</h1>
      <div className="tabs">
        <button onClick={() => setVista("formulario")}>Nueva Muerte</button>
        <button onClick={() => setVista("muertes")}>Ver Muertes</button>
      </div>

      <div className="vista-contenido">
        {vista === "formulario" ? (
          <FormularioMuerte />
        ) : (
          <Suspense fallback={<div>Cargando muertes...</div>}>
            <VerMuertes />
          </Suspense>
        )}
      </div>
    </div>
  );
}

export default App;
