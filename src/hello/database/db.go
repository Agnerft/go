package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello/models"
	"net/http"
)

var clienteConfig []models.ClienteConfig

func AtualizarINUSE(resourceID int) int {
	jsons, _ := BuscaPorDoc(12310400000182)

	// Defina a URL do servidor JSON Server e o JSON de atualização
	serverURL := "https://basesip.makesystem.com.br" // Substitua pela URL correta do seu servidor JSON Server

	updateData := make(map[string]interface{})

	// for i := range clienteConfig[0].QuantRamaisOpen {
	// 	ramalDesejado, _ := strconv.Atoi(clienteConfig[0].Ramal)

	// 	if clienteConfig[0].QuantRamaisOpen[i].Ramal == ramalDesejado {
	// 		clienteConfig[0].QuantRamaisOpen[i].INUSE = true

	// 		break // Parar o loop após encontrar o ramal desejado
	// 	}

	// }

	fmt.Println(jsons)
	fmt.Println(serverURL)

	updateJSON, err := json.Marshal(updateData)
	if err != nil {
		fmt.Println("Erro ao codificar os dados de atualização em JSON:", err)
		//return
	}

	// Crie uma solicitação HTTP PATCH para atualizar o recurso
	url := fmt.Sprintf("%s/clientes/%d", serverURL, resourceID)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(updateJSON))
	if err != nil {
		fmt.Println("Erro ao criar a solicitação HTTP PATCH:", err)
		//return
	}

	req.Header.Set("Content-Type", "application/json")

	// Faça a solicitação HTTP PATCH
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na solicitação HTTP PATCH:", err)
		//return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Recurso atualizado com sucesso.")
	} else {
		fmt.Println("Erro ao atualizar o recurso. Status code:", resp.StatusCode)
	}

	return resourceID

}
