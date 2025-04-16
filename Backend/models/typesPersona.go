package models

type Personas struct {
	Id       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	FotoURL  string `json:"foto_url"`
	CreadoEn string `json:"creado_en"`
}
