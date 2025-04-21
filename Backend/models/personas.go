package models

import (
	"github.com/Santiageoff/Death-Note/bd"
)

// Insertar persona en la base de datos
func CreatePersona(P Personas) error {
	bd, err := bd.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO personas (nombre, apellido, foto_url, causa_muerte) VALUES (?, ?, ?, ?)", P.Nombre, P.Apellido, P.FotoURL, P.CausaMuerte)
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

// Actualizar persona existente
func UpdatePersona(P Personas) error {
	bd, err := bd.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE personas SET nombre = ?, apellido = ?, foto_url = ?, causa_muerte = ? WHERE id = ?", P.Nombre, P.Apellido, P.FotoURL, P.CausaMuerte, P.Id)
	return err
}

// Obtener todas las personas
func GetPersona() ([]Personas, error) {
	personas := []Personas{}
	bd, err := bd.GetDB()
	if err != nil {
		return personas, err
	}
	rows, err := bd.Query("SELECT id, nombre, apellido, foto_url, causa_muerte, fecha_muerte FROM personas")
	if err != nil {
		return personas, err
	}
	for rows.Next() {
		var P Personas
		err = rows.Scan(&P.Id, &P.Nombre, &P.Apellido, &P.FotoURL, &P.CausaMuerte, &P.FechaMuerte)
		if err != nil {
			return personas, err
		}
		personas = append(personas, P)
	}
	return personas, nil
}

// Obtener persona por ID
func GetPersonaById(id int64) (Personas, error) {
	var P Personas
	bd, err := bd.GetDB()
	if err != nil {
		return P, err
	}
	row := bd.QueryRow("SELECT id, nombre, apellido, foto_url, causa_muerte, fecha_muerte FROM personas WHERE id = ?", id)
	err = row.Scan(&P.Id, &P.Nombre, &P.Apellido, &P.FotoURL, &P.CausaMuerte, &P.FechaMuerte)
	return P, err
}
