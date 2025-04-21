import React from 'react';
import { Link } from "react-router-dom";
import "./home.css";

const Home = () => {
  const getBackgroundImage = () => {
    return "/pokemonInicio.jpg"; // Ruta correcta desde "public/"
  };

  return (
    <main
      className="home-container"
      style={{ backgroundImage: `url(${getBackgroundImage()})` }}
    >
      <h1 className="home-title">Bienvenido a la Death Note</h1>
      <Link to="/persona">
        <button className="home-button" aria-label="Entrar a la Death Note">
          Entrar a la Death Note 
        </button>
      </Link>
    </main>
  );
};

export default Home;
