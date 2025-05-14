package models

import (
	"fmt"
	"strings"
	"time"
)

// Mapa para controlar si ya se procesó una persona
var personaProcesada = make(map[string]bool)

func generarClavePersona(p Personas) string {
	return fmt.Sprintf("%s-%s-%s", p.Nombre, p.Apellido, p.FotoURL)
}

func FotoMuerte(personas *Personas) bool {
	if strings.TrimSpace(personas.FotoURL) == "" {
		fmt.Println("No hay foto, no hay muerte")
		return false
	}
	return true
}

func EjecutarMuerte(persona Personas) {
	clave := generarClavePersona(persona)

	if personaProcesada[clave] {
		fmt.Println("La persona ya ha sido procesada, no se ejecutará nuevamente.")
		return
	}

	personaProcesada[clave] = true

	go func(p Personas) {
		if !FotoMuerte(&p) {
			fmt.Println("Persona rechazada por falta de foto")
			return
		}

		if strings.TrimSpace(p.CausaMuerte) == "" {
			p.CausaMuerte = "Ataque al corazón"
			fmt.Println("Causa no especificada. Se usará 'Ataque al corazon'")
		}

		if strings.TrimSpace(p.DetallesMuerte) != "" {
			fmt.Println("Esperando 6 minutos y 15 segundos para ejecutar muerte con detalles...")
			time.Sleep(15 * time.Second)
		} else {
			fmt.Println("Esperando 10 segundos para ejecutar muerte sin detalles...")
			time.Sleep(10 * time.Second)
		}

		err := CreatePersona(p)
		if err != nil {
			fmt.Println("Error registrando la muerte:", err)
		} else {
			fmt.Printf("Persona registrada: %s %s\n", p.Nombre, p.Apellido)
		}
	}(persona)
}
