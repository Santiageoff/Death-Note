import React from 'react'
import { useNavigate } from "react-router-dom";
import "../css/renunciar.css"

function Renunciar() {
  const navigate = useNavigate();

  const continuar = () => {
    navigate("/despedida");
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>¿Seguro que quieres renunciar?</h2>
        <h3>Recuerda:</h3>
        <ul>
          <li>Perderas los recuerdos de la Death Note</li>
          <li>No podras ir ni al cielo ni al infierno</li>
        </ul>
        <div className="modal-buttons">
          <button onClick={continuar}>Aceptar</button>
        </div>
      </div>
    </div>
  );
}

export default Renunciar;
