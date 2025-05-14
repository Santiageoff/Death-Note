package models

import (
	"fmt"
	"time"

	"github.com/Santiageoff/Death-Note/bd"
)

// Insertar persona
func CreatePersona(P Personas) error {

	bd, err := bd.GetDB()
	if err != nil {
		return err
	}

	// Asignar fecha de muerte al momento de la inserción
	P.FechaMuerte = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("Insertando persona:", P)
	_, err = bd.Exec(`
		INSERT INTO personas 
		(nombre, apellido, foto_url, causa_muerte, detalles_muerte, fecha_muerte)
		VALUES (?, ?, ?, ?, ?, ?)`,
		P.Nombre, P.Apellido, P.FotoURL, P.CausaMuerte, P.DetallesMuerte, P.FechaMuerte)
	return err
}

// Eliminar persona por ID
func DeletePersona(id int64) error {
	bd, err := bd.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM personas WHERE id = ?", id)
	return err
}

// Actualizar persona
func UpdatePersona(P Personas) error {
	bd, err := bd.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec(`
		UPDATE personas SET 
		nombre = ?, apellido = ?, foto_url = ?, causa_muerte = ?, detalles_muerte = ?
		WHERE id = ?`,
		P.Nombre, P.Apellido, P.FotoURL, P.CausaMuerte, P.DetallesMuerte, P.Id)
	return err
}

// Obtener todas las personas
func GetPersona() ([]Personas, error) {
	personas := []Personas{}
	bd, err := bd.GetDB()
	if err != nil {
		return personas, err
	}
	rows, err := bd.Query(`
		SELECT id, nombre, apellido, foto_url, causa_muerte, detalles_muerte, fecha_muerte 
		FROM personas`)
	if err != nil {
		return personas, err
	}
	defer rows.Close()

	for rows.Next() {
		var P Personas
		err = rows.Scan(&P.Id, &P.Nombre, &P.Apellido, &P.FotoURL, &P.CausaMuerte, &P.DetallesMuerte, &P.FechaMuerte)
		if err != nil {
			return personas, err
		}
		personas = append(personas, P)
	}
	return personas, nil
}

// Obtener una persona por ID
func GetPersonaById(id int64) (Personas, error) {
	var P Personas
	bd, err := bd.GetDB()
	if err != nil {
		return P, err
	}
	row := bd.QueryRow(`
		SELECT id, nombre, apellido, foto_url, causa_muerte, detalles_muerte, fecha_muerte 
		FROM personas WHERE id = ?`, id)

	err = row.Scan(&P.Id, &P.Nombre, &P.Apellido, &P.FotoURL, &P.CausaMuerte, &P.DetallesMuerte, &P.FechaMuerte)
	return P, err
}
