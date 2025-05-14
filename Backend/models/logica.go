package models

import (
	"fmt"
	"strings"
	"time"
)

// Verifica si la persona tiene una foto.
func FotoMuerte(personas *Personas) {
	if personas.FotoURL == "" {
		fmt.Println("No hay foto, no hay muerte")
		return
	}
}

// causaMuerte simula el proceso de muerte según las reglas del Death Note.
func CausaMuerte(muertes *Personas) {
	// Si no se especifica la causa de la muerte, se asigna "Ataque al corazón".
	if strings.TrimSpace(muertes.CausaMuerte) == "" {
		muertes.CausaMuerte = "Ataque al corazón"
		fmt.Println("Causa no especificada. Se usará 'Ataque al corazón'")
	}

	// Si se detalla la causa de la muerte, esperamos 6 minutos y 40 segundos.
	if muertes.DetallesMuerte != "" {
		fmt.Println("Esperando 6 minutos y 40 segundos antes de especificar detalles de la muerte...")
		time.Sleep(6*time.Minute + 40*time.Second)
	} else {
		// Si no se detalla causa, esperamos 40 segundos.
		fmt.Println("Esperando 40 segundos antes de morir...")
		time.Sleep(40 * time.Second)
	}

	// Ahora simplemente registramos la persona en la base de datos
	err := CreatePersona(*muertes)
	if err != nil {
		fmt.Println("Error registrando la muerte:", err)
	} else {
		fmt.Println("Persona registrada correctamente:", muertes.Nombre, muertes.Apellido)
	}
}
