package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ramal struct {
	Ramal int  `json:"ramal"`
	INUSE bool `json:"INUSE"`
}

type JSONConvertible interface {
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
}

type ClienteTeste struct {
	ID              string  `json:"id"`
	Doc             string  `json:"doc"`
	Cliente         string  `json:"cliente"`
	Porta           string  `json:"porta"`
	GrupoRecurso    string  `json:"grupoRecurso"`
	LinkGvc         string  `json:"linkGvc"`
	Ramal           string  `json:"ramal"`
	Senha           string  `json:"senha"`
	QuantRamaisOpen []Ramal `json:"quantRamaisOpen"`
}

func NewClienteTeste() ClienteTeste {

	return ClienteTeste{
		LinkGvc:         ".gvctelecom.com.br:",
		Senha:           "@abc",
		QuantRamaisOpen: []Ramal{},
	}
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

/*
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

*/

// ToJSON serializa a struct para JSON.
func (c *ClienteTeste) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON desserializa a struct a partir de JSON.
func (c *ClienteTeste) FromJSON(data []byte) error {
	return json.Unmarshal(data, c)
}

func ObterClienteExistente(id int, url string) (*ClienteConfig, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d", url, id))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("Erro ao finalizar a abertura da requisição: %d", resp.StatusCode)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cliente não encontrado, código de status: %d", resp.StatusCode)
	}

	var cliente ClienteConfig
	err = json.NewDecoder(resp.Body).Decode(&cliente)
	if err != nil {
		return nil, err
	}

	return &cliente, nil
}

func AtualizarCliente(cliente *ClienteConfig, url string) error {
	// Serializar a struct para JSON.
	jsonBytes, err := json.Marshal(cliente)
	if err != nil {
		return err
	}

	// Enviar uma requisição HTTP PUT (ou PATCH) com o JSON serializado.
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("Erro ao finalizar a abertura da requisição: %d", resp.StatusCode)
		}

	}(resp.Body)

	// Verificar a resposta da requisição.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("falha ao atualizar o cliente, código de status: %d", resp.StatusCode)
	}

	fmt.Println("Cliente atualizado com sucesso!")
	return nil
}
