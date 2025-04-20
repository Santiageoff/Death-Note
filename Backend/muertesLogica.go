package main

import (
	"fmt"
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
func causaMuerte(muertes *models.Muertes) {
	if muertes.CausaMuerte == "" {
		fmt.Println("Causa no especificada. Se usará 'Ataque al corazón'")
		muertes.CausaMuerte = "Ataque al corazón"
	}

	if muertes.CausaMuerte != "" {
		fmt.Println("Esperando 6 minutos y 40 segundos antes de especificar detalles de la muerte...")
		time.Sleep(6*time.Minute + 40*time.Second)
	} else {
		fmt.Println("Esperando 40 segundos antes de morir...")
		time.Sleep(40 * time.Second)
	}

	muertes.HoraMuerte = time.Now().Format("2006-01-02 15:04:05")
	muertes.EstaMuerto = true

	err := models.RegistrarMuerte(*muertes)
	if err != nil {
		fmt.Println("Error registrando la muerte:", err)
	}
}
