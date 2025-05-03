import React from 'react';
import { useNavigate } from "react-router-dom";
import "../css/despedida.css";

function Despedida() {
  const navigate = useNavigate();

  const continuar = () => {
    navigate("/");
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h1>Has renunciado a la Death Note</h1>
        <p>Tu memoria ha sido borrada, y tu destino ya está sellado.</p>
        <div className="modal-buttons">
          <button className="volver-btn" onClick={continuar}>Volver al inicio</button>
        </div>
      </div>
    </div>
  );
}

export default Despedida;
