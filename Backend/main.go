package main

import (
	"log"
	"net/http"
	"time"

	bd "github.com/Santiageoff/Death-Note/bd"
	routes "github.com/Santiageoff/Death-Note/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Intentar obtener conexión a la base de datos
	bd, err := bd.GetDB()
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
	routes.SetupRoutesForPersonas(router)
	// Servir archivos estáticos desde la carpeta /uploads
	fs := http.FileServer(http.Dir("./uploads"))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

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
