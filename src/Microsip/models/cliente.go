package models

type Ramal struct {
	Ramal int
	InUse bool
}
type ClienteConfig struct {
	ID           int     `json:"id"`
	Doc          int     `json:"doc"`
	Cliente      string  `json:"cliente"`
	GrupoRecurso string  `json:"grupoRecurso"`
	LinkGvc      string  `json:"linkGvc"`
	Porta        string  `json:"porta"`
	Ramal        string  `json:"ramal"`
	Senha        string  `json:"senha"`
	QuantRamais  []Ramal `json:"quantRamaisOpen"`
}

type Cliente struct {
	Clientes []ClienteConfig `json:"clientes"`
}

type Doc []struct {
	Doc string `json: "doc"`
}
