import React, { useState } from "react";

function FormularioMuerte() {
  const [formData, setFormData] = useState({
    nombre: "",
    apellido: "",
    foto_url: "",
    causa_muerte: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const response = await fetch("http://localhost:8080/personas", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    });

    if (response.ok) {
      alert("Persona registrada");
      setFormData({
        nombre: "",
        apellido: "",
        foto_url: "",
        causa_muerte: "",
      });
    } else {
      alert("Error al registrar persona");
    }
  };

  return (
    <form className="formulario" onSubmit={handleSubmit}>
      <input type="text" name="nombre" placeholder="Nombre" value={formData.nombre} onChange={handleChange} required />
      <input type="text" name="apellido" placeholder="Apellido" value={formData.apellido} onChange={handleChange} required />
      <input type="text" name="foto_url" placeholder="URL de foto" value={formData.foto_url} onChange={handleChange} required />
      <input type="text" name="causa_muerte" placeholder="Causa de muerte (opcional)" value={formData.causa_muerte} onChange={handleChange} />
      <button type="submit">Escribir en el Death Note</button>
    </form>
  );
}

export default FormularioMuerte;
