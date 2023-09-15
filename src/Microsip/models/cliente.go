package models

type Ramal struct {
	Ramal int
	InUse bool
}
type ClienteConfig struct {
	ID           int     `json:"id"`
	Doc          string  `json:"doc"`
	Cliente      string  `json:"cliente"`
	QuantRamais  []Ramal `json:"quantRamais"`
	GrupoRecurso string  `json:"grupoRecurso"`
	LinkGvc      string  `json:"linkGvc"`
	Porta        string  `json:"porta"`
	Ramal        string  `json:"ramal"`
	Senha        string  `json:"senha"`
}
