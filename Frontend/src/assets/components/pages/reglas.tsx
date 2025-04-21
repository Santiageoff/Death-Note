import React from 'react'
import { useNavigate } from "react-router-dom";
import "./reglas.css";

function Reglas() {
  const navigate = useNavigate();

  const continuar = () => {
    navigate("/persona");
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>Reglas de la Death Note</h2>
        <ul>
          <li>• La persona cuyo nombre sea escrito en este cuaderno morirá.</li>
          <li>• Si la causa de la muerte no es especificada, la persona morirá de un ataque al corazón.</li>
          <li>• Después de escribir el nombre, el usuario tiene 40 segundos para escribir la causa de muerte.</li>
          <li>• Si la causa de la muerte es especificada, se tienen 6 minutos y 40 segundos adicionales para escribir los detalles específicos.</li>
          <li>• Si el usuario no tiene en mente el rostro de la persona que está escribiendo, la Death Note no funcionará.</li>
          <li>• El cuaderno solo afecta a los humanos.</li>
          <li>• El usuario puede renunciar a la Death Note, perdiendo sus recuerdos sobre ella.</li>
          <li>• Un humano que usa la Death Note no podrá ir ni al cielo ni al infierno.</li>
        </ul>
        <div className="modal-buttons">
          <button onClick={continuar}>Aceptar y continuar</button>
        </div>
      </div>
    </div>
  );
}

export default Reglas;
