package models

import (
	"github.com/Santiageoff/Death-Note/bd"
)

// Insertar muerte en la base de datos
func RegistrarMuerte(m Muertes) error {
	bd, err := bd.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO muertes (person_id, causa_muerte, hora_muerte, esta_muerta) VALUES (?, ?, ?, ?)", m.PersonID, m.CausaMuerte, m.HoraMuerte, m.EstaMuerto)
	return err
}

// Obtener todas las muertes
func GetMuerte() ([]Muertes, error) {
	muertes := []Muertes{}
	bd, err := bd.GetDB()
	if err != nil {
		return muertes, err
	}
	rows, err := bd.Query("SELECT person_id, causa_muerte, hora_muerte, esta_muerta FROM muertes")
	if err != nil {
		return muertes, err
	}
	for rows.Next() {
		var M Muertes
		err = rows.Scan(&M.PersonID, &M.CausaMuerte, &M.HoraMuerte, &M.EstaMuerto)
		if err != nil {
			return muertes, err
		}
		muertes = append(muertes, M)
	}
	return muertes, nil
}

// Obtener muerte por ID
func GetMuerteById(id int64) (Muertes, error) {
	var M Muertes
	bd, err := bd.GetDB()
	if err != nil {
		return M, err
	}
	row := bd.QueryRow("SELECT person_id, causa_muerte, hora_muerte, esta_muerta FROM muertes WHERE id = ?", id)
	err = row.Scan(&M.PersonID, &M.CausaMuerte, &M.HoraMuerte, &M.EstaMuerto)
	return M, err
}
