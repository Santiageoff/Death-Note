package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Intentar obtener conexión a la base de datos
	bd, err := getDB()
	if err != nil {
		log.Printf("Error con la base de datos: " + err.Error())
		return
	} else {
		// Verificar si se puede hacer ping a la base
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al hacer conexión con la base. Verifique credenciales. Error: " + err.Error())
			return
		}
	}

	// Definir rutas HTTP
	router := mux.NewRouter()
	setupRoutesForPersonas(router)

	port := "localhost:8080"

	// Configurar servidor con tiempos de espera
	server := &http.Server{
		Handler:      router,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Servidor iniciado en %s", port)
	log.Fatal(server.ListenAndServe())
}
