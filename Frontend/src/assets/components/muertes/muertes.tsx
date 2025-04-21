// src/components/muertes/muertes.tsx
import React, { useEffect, useState } from 'react'

type Persona = {
  id: number;
  nombre: string;
  apellido: string;
  foto_url: string;
  causa_muerte: string;
  hora_muerte: string;
};

const Muertes = () => {
  const [personas, setPersonas] = useState<Persona[]>([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/personas/muertas')
      .then((res) => res.json())
      .then((data) => setPersonas(data))
      .catch((err) => console.error("Error al obtener las muertes:", err));
  }, []);

  return (
    <div className="muertes">
      <h1>Lista de Muertes</h1>
      <div className="lista">
        {personas.map((p) => (
          <div key={p.id} className="persona">
            <img src={p.foto_url} alt={`Foto de ${p.nombre}`} width={150} />
            <h2>{p.nombre} {p.apellido}</h2>
            <p><strong>Causa:</strong> {p.causa_muerte}</p>
            <p><strong>Hora:</strong> {p.hora_muerte}</p>
          </div>
        ))}
      </div>
    </div>
  )
}

export default Muertes;
