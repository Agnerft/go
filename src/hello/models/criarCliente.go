package models

import (
	"bytes"
	"log"
	"net/http"
)

func CriarNovoCliente(cliente *ClienteTeste, url string) ([]byte, error) {
	// Serializar a struct para JSON.
	jsonBytes, err := cliente.ToJSON()
	if err != nil {
		log.Printf("Erro a serializar: %v", err)
	}

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))

	defer resp.Body.Close()

	return jsonBytes, nil
}

func (c *ClienteConfig) AdicionarRamais(inicio int, quantidade int) (string, error) {

	for i := inicio; i < inicio+quantidade; i++ {
		novoRamal := Ramal{
			Ramal: i,
			INUSE: false,
		}
		c.QuantRamaisOpen = append(c.QuantRamaisOpen, novoRamal)

	}
	stringRamais := "Ramais criados com sucesso."
	return stringRamais, nil

}
