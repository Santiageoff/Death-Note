package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Santiageoff/Death-Note/bd"
	"github.com/Santiageoff/Death-Note/models"
	"github.com/Santiageoff/Death-Note/utils"
	"github.com/gorilla/mux"
)

// Configuración de rutas para el manejo de personas
func SetupRoutesForPersonas(router *mux.Router) {
	// Habilitar CORS
	enableCORS(router)

	// Ruta GET: lista todas las personas
	router.HandleFunc("/personas", func(w http.ResponseWriter, r *http.Request) {
		personas, err := models.GetPersona()
		if err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(personas, w)
	}).Methods(http.MethodGet)

	// Ruta GET con ID: obtener una sola persona
	router.HandleFunc("/personas/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := utils.StringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		persona, err := models.GetPersonaById(id)
		if err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(persona, w)
	}).Methods(http.MethodGet)

	// Ruta POST: crear nueva persona con lógica Death Note
	// Ruta POST: crear nueva persona con lógica Death Note
	router.HandleFunc("/personas", func(w http.ResponseWriter, r *http.Request) {
		var persona models.Personas
		if err := json.NewDecoder(r.Body).Decode(&persona); err != nil {
			http.Error(w, "Datos inválidos en la solicitud", http.StatusBadRequest)
			return
		}

		// Verificar si hay foto. Si no, no se registra.
		if !models.FotoMuerte(&persona) {
			http.Error(w, "No se puede registrar sin foto", http.StatusBadRequest)
			return
		}

		// Ejecutar la lógica de muerte en una goroutine (asincronía)
		go models.EjecutarMuerte(persona)

		// Responder de inmediato al frontend
		respondWithSuccess(map[string]string{
			"status":  "Muerte en proceso",
			"message": "La persona será registrada pronto... si es que no sobrevive.",
		}, w)
	}).Methods(http.MethodPost)

	// Ruta POST: subir imagen y devolver la ruta local
	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "No se pudo procesar la imagen", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("foto")
		if err != nil {
			http.Error(w, "No se encontró el archivo", http.StatusBadRequest)
			return
		}
		defer file.Close()

		uploadDir := "./uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			os.Mkdir(uploadDir, os.ModePerm)
		}

		dst := filepath.Join(uploadDir, handler.Filename)
		out, err := os.Create(dst)
		if err != nil {
			http.Error(w, "No se pudo guardar la imagen", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, "Error al guardar la imagen", http.StatusInternalServerError)
			return
		}

		relativePath := "/uploads/" + handler.Filename
		json.NewEncoder(w).Encode(map[string]string{"foto_url": relativePath})
	}).Methods(http.MethodPost)

	// Ruta DELETE: eliminar persona por ID
	router.HandleFunc("/personas/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := utils.StringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		if err := models.DeletePersona(id); err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)

	// Ruta DELETE: eliminar todas las personas
	router.HandleFunc("/renunciar", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Llamar a la función para eliminar todas las personas
		err := models.DeleteAllPersonas()
		if err != nil {
			respondWithError(err, w)
			return
		}

		// Responder al cliente que la base de datos fue limpiada
		respondWithSuccess(map[string]string{
			"status":  "Éxito",
			"message": "La base de datos ha sido limpiada exitosamente.",
		}, w)
	}).Methods(http.MethodDelete)

	// Ruta PUT: actualizar persona por ID
	router.HandleFunc("/personas/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := utils.StringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}

		var persona models.Personas
		if err := json.NewDecoder(r.Body).Decode(&persona); err != nil {
			http.Error(w, "Datos inválidos en la solicitud", http.StatusBadRequest)
			return
		}
		persona.Id = id

		if err := models.UpdatePersona(persona); err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodPut)

	// Ruta para servir el favicon
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}

// CORS
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", bd.AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", bd.AllowedCORSDomain)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, req)
	})
}

// Manejo de errores y respuestas
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
