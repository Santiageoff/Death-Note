package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Configuración de rutas para el manejo de personas
func setupRoutesForPersonas(router *mux.Router) {
	// Habilitar CORS
	enableCORS(router)

	// Ruta GET: lista todas las personas
	router.HandleFunc("/personas", func(w http.ResponseWriter, r *http.Request) {
		personas, err := getPersona()
		if err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(personas, w)
	}).Methods(http.MethodGet)

	// Ruta GET con ID: obtener una sola persona
	router.HandleFunc("/personas/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		persona, err := getPersonaById(id)
		if err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(persona, w)
	}).Methods(http.MethodGet)

	// Ruta POST: crear nueva persona
	router.HandleFunc("/personas", func(w http.ResponseWriter, r *http.Request) {
		var persona Personas
		err := json.NewDecoder(r.Body).Decode(&persona)
		if err != nil {
			respondWithError(err, w)
			return
		}
		if err := createPersona(persona); err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(true, w)
	}).Methods(http.MethodPost)

	// Ruta PUT: actualizar persona
	router.HandleFunc("/personas", func(w http.ResponseWriter, r *http.Request) {
		var persona Personas
		err := json.NewDecoder(r.Body).Decode(&persona)
		if err != nil {
			respondWithError(err, w)
			return
		}
		if err := updatePersona(persona); err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(true, w)
	}).Methods(http.MethodPut)

	// Ruta DELETE: eliminar persona
	router.HandleFunc("/personas/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		if err := deletePersona(id); err != nil {
			respondWithError(err, w)
			return
		}
		respondWithSuccess(true, w)
	}).Methods(http.MethodDelete)

	// Ruta para servir el favicon
	// Ignorar la solicitud para favicon.ico
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

}

// Configuración general de CORS
func enableCORS(router *mux.Router) {
	// Permitir solicitudes de otros orígenes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	// Middleware que aplica los headers CORS a todas las rutas
	router.Use(middlewareCors)
}

// Middleware para CORS
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
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
