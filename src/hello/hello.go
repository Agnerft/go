package main

import (
	"encoding/json"
	"fmt"
	"hello/database"
	"hello/models"
	"hello/utils"
	"os"
	"os/exec"
)

func main() {

	desktopPath, _ := os.UserHomeDir()
	nomeInstalador := "MicroSIP"
	zipURL := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	versao := "-3.21.3"
	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	resultadoIni := utils.SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")
	// Pegando a busca por doc e setando num jsonfile
	jsonfile, _ := database.BuscaPorDoc(44764891000128)

	var clienteConfig []models.ClienteConfig
	if err := json.Unmarshal(jsonfile, &clienteConfig); err != nil {
		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
		return
	}

	resultadoInstalador := utils.SalvarArquivo(zipURL, desktopPath, nomeInstalador, versao+".exe")

	//Executar o instalador do MicroSIP

	cmd := exec.Command(resultadoInstalador, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o instalador:%s ", err)
		return
	}

	fmt.Println("Executou ?")

	// Edição do arquivo .ini
	for _, config := range clienteConfig {
		fmt.Println("Oi")
		linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
		utils.Editor(resultadoIni, 2, "label="+config.Ramal)
		utils.Editor(resultadoIni, 3, "server="+linkCompleto)
		utils.Editor(resultadoIni, 4, "proxy="+linkCompleto)
		utils.Editor(resultadoIni, 5, "domain="+linkCompleto)
		utils.Editor(resultadoIni, 6, "username="+config.Ramal)
		utils.Editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
		utils.Editor(resultadoIni, 8, "authID="+config.Ramal)
	}

	processName := "MicroSIP.exe" // Substitua pelo nome do processo que você deseja encerrar

	cmd3 := exec.Command("taskkill", "/F", "/IM", processName)

	// Redirecionar saída e erro, se necessário
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr

	err := cmd3.Run()
	if err != nil {
		fmt.Println("Erro ao executar o comando:", err)
		return
	}
	fmt.Println("Processo encerrado com sucesso.")
	// fmt.Println(clienteConfig[len(clienteConfig)-1])
	// fmt.Println("passou aqui?")

	//novoCliente := &clienteConfig[len(clienteConfig)-1]

	// fmt.Printf(novoCliente.Cliente)

	// novoCliente.AdicionarRamais(7800, 10)

	tests := &models.ClienteConfig{Doc: 37986991000133, Cliente: "conectasul",
		GrupoRecurso:    "conectasul",
		LinkGvc:         ".gvctelecom.com.br:",
		Porta:           "4177",
		Ramal:           "7801",
		Senha:           "@abc",
		QuantRamaisOpen: []models.Ramal{},
	}
	tests.AdicionarRamais(7801, 10)

	models.CriarNovoCliente(tests, "http://localhost:3004/clientes")

	// id := database.AtualizarINUSE(1)
	// fmt.Println(id)
	// quantRamaisOpen := map[string]interface{}
	// clientes := quantRamaisOpen["clientes"].([]map[string]interface{})
	// primeiroCliente := clientes[0]
	// fmt.Println("ID do primeiro cliente:", primeiroCliente["id"])

	//O ramal que você deseja encontrar

	// for i := range clienteConfig[0].QuantRamaisOpen {
	// 	ramalDesejado, _ := strconv.Atoi(clienteConfig[0].Ramal)

	// 	if clienteConfig[0].QuantRamaisOpen[i].Ramal == ramalDesejado {
	// 		clienteConfig[0].QuantRamaisOpen[i].INUSE = true

	// 		break // Parar o loop após encontrar o ramal desejado
	// 	}

	// }
	// quantRamaisOpen := make(map[string]interface{})

	// clientes := quantRamaisOpen["clientes"].([]map[string]interface{})
	// primeiroCliente := clientes[0]
	// fmt.Println("ID do primeiro cliente:", primeiroCliente["id"])

	// // Exemplo de como acessar valores em quantRamaisOpen
	// quantRamais := primeiroCliente["quantRamaisOpen"].([]map[string]interface{})
	// fmt.Println("Estado do ramal 7801:", quantRamais[0]["INUSE"])

}
