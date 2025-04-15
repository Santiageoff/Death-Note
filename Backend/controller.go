package main

// Insertar persona en la base de datos
func createPersona(P Personas) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO personas (nombre, apellido, foto_url) VALUES (?, ?, ?)", P.Nombre, P.Apellido, P.FotoURL)
	return err
}

// Eliminar persona por ID
func deletePersona(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM personas WHERE id = ?", id)
	return err
}

// Actualizar persona existente
func updatePersona(Persona Personas) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE personas SET nombre = ?, apellido = ?, foto_url = ? WHERE id = ?", Persona.Nombre, Persona.Apellido, Persona.FotoURL, Persona.Id)
	return err
}

// Obtener todas las personas
func getPersona() ([]Personas, error) {
	personas := []Personas{}
	bd, err := getDB()
	if err != nil {
		return personas, err
	}
	rows, err := bd.Query("SELECT id, nombre, apellido, foto_url, creado_en FROM personas")
	if err != nil {
		return personas, err
	}
	for rows.Next() {
		var P Personas
		err = rows.Scan(&P.Id, &P.Nombre, &P.Apellido, &P.FotoURL, &P.CreadoEn)
		if err != nil {
			return personas, err
		}
		personas = append(personas, P)
	}
	return personas, nil
}

// Obtener persona por ID
func getPersonaById(id int64) (Personas, error) {
	var P Personas
	bd, err := getDB()
	if err != nil {
		return P, err
	}
	row := bd.QueryRow("SELECT id, nombre, apellido, foto_url, creado_en FROM personas WHERE id = ?", id)
	err = row.Scan(&P.Id, &P.Nombre, &P.Apellido, &P.FotoURL, &P.CreadoEn)
	return P, err
}
