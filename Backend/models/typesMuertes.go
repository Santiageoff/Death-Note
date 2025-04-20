package models

type Muertes struct {
	PersonID    int64  `json:"person_id"`
	CausaMuerte string `json:"causa_muerte"`
	HoraMuerte  string `json:"hora_muerte"`
	EstaMuerto  bool   `json:"esta_muerta"`
}
