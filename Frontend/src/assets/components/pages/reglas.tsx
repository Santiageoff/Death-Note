import React from 'react'
import { useNavigate } from "react-router-dom";
import "../css/reglas.css"

function Reglas() {
  const navigate = useNavigate();

  const continuar = () => {
    navigate("/persona");
  };

  return (
    <div className="modal">
      <div className="modal-content">
        <h2>Death Note</h2>
        <h3>Como se usa</h3>
        <ul>
          <li>• La persona cuyo nombre sea escrito en este cuaderno morira.</li>
          <li>• Si la causa de la muerte no es especificada, la persona morira de un ataque al corazon.</li>
          <li>• Despues de escribir el nombre, tiene 40 segundos para escribir la causa de muerte.</li>
          <li>• Si la causa de la muerte es especificada, se tienen 6 minutos y 40 segundos adicionales para escribir los detalles específicos.</li>
          <li>• Si no se tiene en mente el rostro de la persona que está escribiendo, la Death Note no funcionará.</li>
          <li>• El cuaderno solo afecta a los humanos.</li>
          <li>• Se puede renunciar a la Death Note, perdiendo sus recuerdos sobre ella.</li>
          <li>• Un humano que usa la Death Note no podra ir ni al cielo ni al infierno.</li>
        </ul>
        <div className="modal-buttons">
          <button onClick={continuar}>Aceptar y continuar</button>
        </div>
      </div>
    </div>
  );
}

export default Reglas;
