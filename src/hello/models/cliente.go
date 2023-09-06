package models

import "encoding/json"

type Ramal struct {
	Ramal int  `json:"ramal"`
	INUSE bool `json:"INUSE"`
}

type JSONConvertible interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}
type ClienteConfig struct {
	ID              int     `json:"id"`
	Doc             int64   `json:"doc"`
	Cliente         string  `json:"cliente"`
	GrupoRecurso    string  `json:"grupoRecurso"`
	LinkGvc         string  `json:"linkGvc"`
	Porta           string  `json:"porta"`
	Ramal           string  `json:"ramal"`
	Senha           string  `json:"senha"`
	QuantRamaisOpen []Ramal `json:"quantRamaisOpen"`
}

type Cliente struct {
	Clientes []ClienteConfig `json:"clientes"`
}

func AdicionarCliente(mapa map[string]interface{}, id int, doc int, cliente string, grupoRecurso string, linkGvc string, porta string, ramal interface{}, senha string) {

	clientes := []map[string]interface{}{
		{
			"id":              id,
			"doc":             doc,
			"cliente":         cliente,
			"grupoRecurso":    grupoRecurso,
			"linkGvc":         linkGvc,
			"porta":           porta,
			"ramal":           ramal,
			"senha":           senha,
			"quantRamaisOpen": []map[string]interface{}{},
		},
	}

	mapa["clientes"] = clientes
}

func AdicionarRamal(mapa map[string]interface{}, clienteID int, ramalNum int, inUse bool) {
	clientes := mapa["clientes"].([]map[string]interface{})
	for i := range clientes {
		if clientes[i]["id"].(int) == clienteID {
			ramais := clientes[i]["quantRamaisOpen"].([]map[string]interface{})
			ramal := map[string]interface{}{
				"ramal": 780 + ramalNum,
				"INUSE": inUse,
			}
			clientes[i]["quantRamaisOpen"] = append(ramais, ramal)
			break
		}
	}
}

// ToJSON serializa a struct para JSON.
func (c *ClienteConfig) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON desserializa a struct a partir de JSON.
func (c *ClienteConfig) FromJSON(data []byte) error {
	return json.Unmarshal(data, c)
}

// AdicionarRamais adiciona a quantidade especificada de ramais ao slice QuantRamaisOpen.
