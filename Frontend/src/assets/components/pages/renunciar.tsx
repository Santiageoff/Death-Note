import React from 'react';
import { useNavigate } from "react-router-dom";
import "../css/renunciar.css";

function Renunciar() {
  const navigate = useNavigate();

  // Función para manejar la renuncia
  const continuar = async () => {
    try {
      // Realizar la solicitud para borrar los recuerdos en el backend
      const response = await fetch("http://localhost:8080/renunciar", {
        method: "DELETE",
      });

      // Verificar si la respuesta es correcta
      if (!response.ok) {
        alert("Ocurrió un error al borrar los recuerdos.");
        return;
      }

      // Redirigir al usuario a la página de despedida
      setTimeout(() => {
        navigate("/despedida");
      }, 1000);
      
    } catch (error) {
      alert("No se pudo conectar con el servidor.");
      console.error(error);
    }
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>¿Seguro que quieres renunciar?</h2>
        <h3>Recuerda:</h3>
        <ul>
          <li>Perderás los recuerdos de la Death Note</li>
          <li>No podrás ir ni al cielo ni al infierno</li>
        </ul>
        <div className="modal-buttons">
          <button onClick={continuar}>Aceptar</button>
        </div>
      </div>
    </div>
  );
}

export default Renunciar;
