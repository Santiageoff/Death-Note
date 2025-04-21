package routes

import (
	"encoding/json"
	"net/http"

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

	// Ruta POST: crear nueva persona
	router.HandleFunc("/personas", func(w http.ResponseWriter, r *http.Request) {
		var persona models.Personas
		if err := json.NewDecoder(r.Body).Decode(&persona); err != nil {
			http.Error(w, "Datos inválidos en la solicitud", http.StatusBadRequest)
			return
		}
		if err := models.CreatePersona(persona); err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
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
		persona.Id = id // Asegurar que el ID venga de la URL

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

// Configuración general de CORS
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", bd.AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	router.Use(middlewareCors)
}

// Middleware para CORS
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", bd.AllowedCORSDomain)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, req)
	})
}

// Respuesta con error
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

// Respuesta con éxito
func respondWithSuccess(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
