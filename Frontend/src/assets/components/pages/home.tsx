import React from 'react';
import { Link } from "react-router-dom";
import "../css/home.css";

const Home = () => {
  const getBackgroundImage = () => {
    return "/Fondo.jpg";
  };

  return (
    <main
      className="home-container"
      style={{ backgroundImage: `url(${getBackgroundImage()})` }}
    >
      <h1 className="home-title">Death Note</h1>
      <Link to="/reglas">
        <button className="home-button" aria-label="Entrar a la Death Note">
          Entrar a la Death Note 
        </button>
      </Link>
    </main>
  );
};

export default Home;
