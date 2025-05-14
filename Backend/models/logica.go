package models

import (
	"fmt"
	"strings"
	"time"
)

// Mapa en memoria para controlar si la persona ya fue procesada
var personaProcesada = make(map[int64]bool) // ID de persona -> procesada (true/false)

// Verifica si la persona tiene una foto.
func FotoMuerte(personas *Personas) bool {
	if strings.TrimSpace(personas.FotoURL) == "" {
		fmt.Println("No hay foto, no hay muerte")
		return false
	}
	return true
}

// Inicia el proceso de "matar" a una persona según las reglas del Death Note.
// Esta función se ejecuta en una goroutine, por lo que no bloquea la petición HTTP.
func EjecutarMuerte(persona Personas) {
	// Verificar si la persona ya ha sido procesada usando el mapa
	if personaProcesada[persona.Id] {
		fmt.Println("La persona ya ha sido procesada, no se ejecutará nuevamente.")
		return
	}

	// Marcar la persona como procesada en el mapa
	personaProcesada[persona.Id] = true

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
			// Reducir el tiempo a 5 segundos para demostración
			fmt.Println("Esperando 6 minutos y 15 segundos para ejecutar muerte con detalles...")
			time.Sleep( /*6*time.Minute*/ 15 * time.Second)
		} else {
			// Reducir el tiempo a 2 segundos para demostración
			fmt.Println("Esperando 10 segundos para ejecutar muerte sin detalles...")
			time.Sleep(10 * time.Second)
		}

		// Ejecutar el proceso de registro después de la "muerte"
		err := CreatePersona(p)
		if err != nil {
			fmt.Println("Error registrando la muerte:", err)
		} else {
			fmt.Printf("Persona registrada: %s %s\n", p.Nombre, p.Apellido)
		}
	}(persona)
}
