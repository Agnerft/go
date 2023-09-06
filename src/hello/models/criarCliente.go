package models

import (
	"bytes"
	"log"
	"net/http"
)

//var clienteConfig []ClienteConfig

func CriarNovoCliente(cliente *ClienteConfig, url string) ([]byte, error) {
	// Serializar a struct para JSON.
	jsonBytes, err := cliente.ToJSON()
	if err != nil {
		log.Printf("Erro a serializar: %v", err)
	}

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

	defer resp.Body.Close()

	return jsonBytes, nil
}

func (c *ClienteConfig) AdicionarRamais(inicio, quantidade int) {
	for i := inicio; i < inicio+quantidade; i++ {
		novoRamal := Ramal{
			Ramal: i,
			INUSE: false,
		}
		c.QuantRamaisOpen = append(c.QuantRamaisOpen, novoRamal)
	}
}
