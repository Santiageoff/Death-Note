import React, { Suspense, useState, useEffect } from "react";
import { Link } from "react-router-dom";
import FormularioMuerte from "./components/pages/formulario";
import "./App.css";

interface Persona {
  id: number;
  nombre: string;
  apellido: string;
  foto_url: string;
  causa_muerte: string;
  detalles_muerte: string;
  fecha_creacion: string;
  fecha_muerte: string;
}

function VerMuertes() {
  const [data, setData] = useState<Persona[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:8080/personas");
      const result = await response.json();
      console.log("Datos recibidos del backend:");
      console.log(result); // <--- Aquí ves todos los datos recibidos
      setData(result);
    };
    fetchData();
  }, []);

  return (
    <div className="muertes">
      <h1>Lista de Muertes</h1>
      <div className="lista">
        {data.map((user) => {
          const encodedUrl = encodeURIComponent(user.foto_url);
          console.log("Foto original:", user.foto_url);
          console.log("Nombre:", user.nombre);
          console.log("Apellido:", user.apellido);
          console.log("Causa:", user.causa_muerte);



          return (
            <div key={user.id} className="persona">
              <img
                src={`http://localhost:8080${user.foto_url}`}
                alt={`${user.nombre} ${user.apellido}`}
                width={150}
              />
              <h2>{user.nombre} {user.apellido}</h2>
              <p><strong>Causa:</strong> {user.causa_muerte}</p>
              <p><strong>Detalles</strong> {user.detalles_muerte}</p>
              <p><strong>Fecha de Registro:</strong> {user.fecha_creacion}</p>
              <p><strong>Fecha de Muerte:</strong> {user.fecha_muerte}</p>
            </div>
          );
        })}
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
        <Link to="/renunciar">
          <button>Renunciar</button>
        </Link>
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
