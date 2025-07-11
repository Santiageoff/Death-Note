package models

type Personas struct {
	Id             int64  `json:"id"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	FotoURL        string `json:"foto_url"`
	CausaMuerte    string `json:"causa_muerte"`
	DetallesMuerte string `json:"detalles_muerte"`
	FechaCreacion  string `json:"fecha_creacion"`
	FechaMuerte    string `json:"fecha_muerte"`
}
