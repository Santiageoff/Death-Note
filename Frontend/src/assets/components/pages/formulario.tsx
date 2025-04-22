import React, { useState } from "react";

function FormularioMuerte() {
  const [formData, setFormData] = useState({
    nombre: "",
    apellido: "",
    foto_url: "",
    causa_muerte: "",
  });

  const [fotoFile, setFotoFile] = useState<File | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, files } = e.target;

    if (name === "foto_url" && files && files[0]) {
      setFotoFile(files[0]);
    } else {
      setFormData({ ...formData, [name]: value });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // 1. Convertir la imagen a base64 si existe
    let imageUrl = "";
    if (fotoFile) {
      const reader = new FileReader();

      reader.onloadend = () => {
        imageUrl = reader.result as string;

        // 2. Enviar datos al backend
        const finalData = {
          ...formData,
          foto_url: imageUrl,
        };

        fetch("http://localhost:8080/personas", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(finalData),
        })
          .then((response) => {
            if (response.ok) {
              alert("Persona registrada");
              setFormData({
                nombre: "",
                apellido: "",
                foto_url: "",
                causa_muerte: "",
              });
              setFotoFile(null);
            } else {
              alert("Error al registrar persona");
            }
          })
          .catch((error) => {
            alert("Hubo un error con la conexión al servidor.");
            console.error(error);
          });
      };

      reader.readAsDataURL(fotoFile); // Convierte la imagen a base64
    }
  };

  return (
    <form className="formulario" onSubmit={handleSubmit}>
      <div className="foto-container">
      <label htmlFor="foto_url">Sube una foto de la persona</label>
      <input
        type="file"
        id="foto_url"
        name="foto_url"
        accept="image/*"
        onChange={handleChange}
        required
      />
      </div>
      <input
        type="text"
        name="nombre"
        placeholder="Nombre"
        value={formData.nombre}
        onChange={handleChange}
        required
      />
      <input
        type="text"
        name="apellido"
        placeholder="Apellido"
        value={formData.apellido}
        onChange={handleChange}
        required
      />
      <input
        type="text"
        name="causa_muerte"
        placeholder="Causa de muerte (opcional)"
        value={formData.causa_muerte}
        onChange={handleChange}
      />
      <button type="submit">Escribir en el Death Note</button>
    </form>
  );
}

export default FormularioMuerte;
