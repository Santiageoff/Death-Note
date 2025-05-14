import React, { useState, useEffect } from "react";

function FormularioMuerte() {
  const [formData, setFormData] = useState({
    nombre: "",
    apellido: "",
    foto_url: "",
    causa_muerte: "",
    fecha_muerte: "",
  });

  const [fotoFile, setFotoFile] = useState<File | null>(null);
  const [timer, setTimer] = useState<number | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, files } = e.target;

    if (name === "foto_url" && files && files[0]) {
      setFotoFile(files[0]);
    } else {
      setFormData({ ...formData, [name]: value });
    }
  };

  useEffect(() => {
    if (formData.nombre) {
      const timeout = setTimeout(() => {
        if (!formData.causa_muerte) {
          const deathTime = new Date();
          deathTime.setSeconds(deathTime.getSeconds() + 40);  // 40 segundos para morir
          setFormData({ ...formData, fecha_muerte: deathTime.toISOString() });
        }
      }, 40000); // 40 segundos para morir si no hay causa

      setTimer(timeout);
      return () => clearTimeout(timeout); // Limpiar el temporizador cuando el nombre cambie
    }
  }, [formData.nombre]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    let imageUrl = "";
    if (fotoFile) {
      const formData = new FormData();
      formData.append("foto", fotoFile);

      const uploadResponse = await fetch("http://localhost:8080/upload", {
        method: "POST",
        body: formData,
      });

      if (!uploadResponse.ok) {
        alert("Error al cargar la imagen");
        return;
      }

      const uploadData = await uploadResponse.json();
      imageUrl = uploadData.foto_url;
    }

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
            fecha_muerte: "",
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
