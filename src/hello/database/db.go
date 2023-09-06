package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello/models"
	"net/http"
)

var clienteConfig []models.ClienteConfig

<<<<<<< HEAD
=======
func BuscaPorDoc(doc int) (string, error) {
	//doc := 12310400000182
	// ajustar porta
	url := "http://localhost:3004/clientes?doc=" + strconv.Itoa(doc)
	method := "GET"

	//fmt.Println(url)

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return

	}
	//fmt.Println(string(body))

	return string(body), nil
}

>>>>>>> d953a134d8df91c2235d8aa7204028e00eb8a35f
func AtualizarINUSE(resourceID int) int {
	jsons, _ := BuscaPorDoc(12310400000182)

	// Defina a URL do servidor JSON Server e o JSON de atualização
	serverURL := "http://localhost:3004" // Substitua pela URL correta do seu servidor JSON Server

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

<<<<<<< HEAD
=======
}

func CriarCliente(resourceID int) int {
	serverURL := "http://localhost:3004" // Substitua pela URL correta do seu servidor JSON Server

	updateData := make(map[string]interface{})
	models.AdicionarCliente(updateData, 1, 12310400000182, "make2", "make", ".gvctelecom.com.br:", "5071", "7801", "@abc")

	clientes := updateData["clientes"].([]map[string]interface{})
	primeiroCliente := clientes[0]
	quantRamais := primeiroCliente["quantRamaisOpen"].([]map[string]interface{})

	models.AdicionarCliente(updateData, 1, 12310400000182, "make2", "make", ".gvctelecom.com.br:", "5071", "7801", "@abc")

	models.AdicionarRamal(updateData, 1, 1, false)
	models.AdicionarRamal(updateData, 1, 2, false)

	fmt.Println(primeiroCliente)
	fmt.Println(quantRamais)

	updateJSON, err := json.Marshal(primeiroCliente)
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
>>>>>>> d953a134d8df91c2235d8aa7204028e00eb8a35f
}
