package models

type Ramal struct {
	Ramal int
	InUse bool
}
type ClienteConfig struct {
	ID           int     `json:"id"`
	Doc          int     `json:"doc"`
	Cliente      string  `json:"cliente"`
	QuantRamais  []Ramal `json:"quantRamaisOpen"`
	GrupoRecurso string  `json:"grupoRecurso"`
	LinkGvc      string  `json:"linkGvc"`
	Porta        string  `json:"porta"`
	Ramal        string  `json:"ramal"`
	Senha        string  `json:"senha"`
}
