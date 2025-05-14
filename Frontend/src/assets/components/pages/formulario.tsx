import React, { useState } from "react";

interface FormDataType {
  nombre: string;
  apellido: string;
  foto_url: string;
  causa_muerte: string;
  detalles_muerte: string;
}

function FormularioMuerte() {
  const [formData, setFormData] = useState<FormDataType>({
    nombre: "",
    apellido: "",
    foto_url: "",
    causa_muerte: "",
    detalles_muerte: "",
  });

  const [fotoFile, setFotoFile] = useState<File | null>(null);
  const [isWaiting, setIsWaiting] = useState<boolean>(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, files } = e.target;

    if (name === "foto_url" && files && files[0]) {
      setFotoFile(files[0]);
    } else {
      setFormData((prev) => ({ ...prev, [name]: value }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
  e.preventDefault();

  const fileInput = document.getElementById("foto_url") as HTMLInputElement;
  const selectedFile = fileInput?.files?.[0];

  if (!selectedFile) {
    alert("Debes subir una foto antes de continuar.");
    return;
  }

  setIsWaiting(true);

  try {
    const formDataUpload = new FormData();
    formDataUpload.append("foto", selectedFile);

    const uploadResponse = await fetch("http://localhost:8080/upload", {
      method: "POST",
      body: formDataUpload,
    });

    if (!uploadResponse.ok) {
      alert("Error al cargar la imagen");
      setIsWaiting(false);
      return;
    }

    const uploadData = await uploadResponse.json();
    const imageUrl: string = uploadData.foto_url;

    const finalData = {
      nombre: formData.nombre.trim(),
      apellido: formData.apellido.trim(),
      foto_url: imageUrl,
      causa_muerte: formData.causa_muerte.trim(),
      detalles_muerte: formData.detalles_muerte.trim(),
    };

    const response = await fetch("http://localhost:8080/personas", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(finalData),
    });

    if (response.ok) {
      alert("La muerte ha sido registrada en las sombras...");
      setFormData({
        nombre: "",
        apellido: "",
        foto_url: "",
        causa_muerte: "",
        detalles_muerte: "",
      });
      fileInput.value = ""; // Limpiar el input de imagen
    } else {
      alert("Error al registrar la persona.");
    }
  } catch (error) {
    alert("Error en la conexión al servidor.");
    console.error(error);
  } finally {
    setIsWaiting(false);
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
          disabled={isWaiting}
        />
      </div>
      <input
        type="text"
        name="nombre"
        placeholder="Nombre"
        value={formData.nombre}
        onChange={handleChange}
        required
        disabled={isWaiting}
      />
      <input
        type="text"
        name="apellido"
        placeholder="Apellido"
        value={formData.apellido}
        onChange={handleChange}
        required
        disabled={isWaiting}
      />
      <input
        type="text"
        name="causa_muerte"
        placeholder="Causa de muerte (opcional)"
        value={formData.causa_muerte}
        onChange={handleChange}
        disabled={isWaiting}
      />
      <input
        type="text"
        name="detalles_muerte"
        placeholder="Detalles de la muerte (opcional)"
        value={formData.detalles_muerte}
        onChange={handleChange}
        disabled={isWaiting}
      />
      <button type="submit" disabled={isWaiting}>
        Escribir en el Death Note
      </button>
    </form>
  );
}

export default FormularioMuerte;
