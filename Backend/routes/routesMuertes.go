package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Santiageoff/Death-Note/bd"
	"github.com/Santiageoff/Death-Note/models"
	"github.com/Santiageoff/Death-Note/utils"
	"github.com/gorilla/mux"
)

// Configuración de rutas para el manejo de muertes
func SetupRoutesForMuertes(router *mux.Router) {
	// Habilitar CORS
	enableCORSM(router)

	// Ruta GET: lista todas las muertes
	router.HandleFunc("/muertes", func(w http.ResponseWriter, r *http.Request) {
		muertes, err := models.GetMuerte()
		if err != nil {
			respondWithErrorM(err, w)
			return
		}
		respondWithSuccessM(muertes, w)
	}).Methods(http.MethodGet)

	// Ruta GET con ID: obtener una sola muerte
	router.HandleFunc("/muertes/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := utils.StringToInt64(idAsString)
		if err != nil {
			respondWithErrorM(err, w)
			return
		}
		muerte, err := models.GetMuerteById(id)
		if err != nil {
			respondWithErrorM(err, w)
			return
		}
		respondWithSuccessM(muerte, w)
	}).Methods(http.MethodGet)

	// Ruta para servir el favicon (opcional)
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}

// Configuración general de CORS
func enableCORSM(router *mux.Router) {
	// Permitir solicitudes de otros orígenes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", bd.AllowedCORSDomain)
	}).Methods(http.MethodOptions)

	// Middleware que aplica los headers CORS a todas las rutas
	router.Use(middlewareCorsM)
}

// Middleware para CORS
func middlewareCorsM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", bd.AllowedCORSDomain)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, req)
	})
}

// Respuesta con error
func respondWithErrorM(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

// Respuesta con éxito
func respondWithSuccessM(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
