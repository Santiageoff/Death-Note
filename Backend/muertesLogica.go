package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Santiageoff/Death-Note/models"
)

// Verifica si la persona tiene una foto.
func FotoMuerte(personas *models.Personas) {
	if personas.FotoURL == "" {
		fmt.Println("No hay foto, no hay muerte")
		return
	}
}

// causaMuerte simula el proceso de muerte según las reglas del Death Note.
func causaMuerte(muertes *models.Personas) {
	// Si no se especifica la causa de la muerte, se asigna "Ataque al corazón".
	if strings.TrimSpace(muertes.CausaMuerte) == "" {
		muertes.CausaMuerte = "Ataque al corazón"
		fmt.Println("Causa no especificada. Se usará 'Ataque al corazón'")
	}

	// Si se especifica la causa de la muerte, esperamos 6 minutos y 40 segundos.
	if muertes.CausaMuerte != "" {
		fmt.Println("Esperando 6 minutos y 40 segundos antes de especificar detalles de la muerte...")
		time.Sleep(6*time.Minute + 40*time.Second)
	} else {
		// Si no se especifica causa, esperamos 40 segundos.
		fmt.Println("Esperando 40 segundos antes de morir...")
		time.Sleep(40 * time.Second)
	}

	// Ahora simplemente registramos la persona en la base de datos
	err := models.CreatePersona(*muertes)
	if err != nil {
		fmt.Println("Error registrando la muerte:", err)
	} else {
		fmt.Println("Persona registrada correctamente:", muertes.Nombre, muertes.Apellido)
	}
}
