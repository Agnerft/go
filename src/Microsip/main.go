package main

import (
	"encoding/json"
	"fmt"
	"github.com.br/agnerft/database"
	"github.com.br/agnerft/models"
	"github.com.br/agnerft/utils"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

var clientConfig []models.ClienteConfig
var docCliente models.Doc

func main() {

	versao := "-3.21.3"
	// URL do arquivo ZIP
	zipURL := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	nomeInstalador := "MicroSIP"

	desktopPath, _ := os.UserHomeDir()

	var doc string

	for {
		fmt.Print("Digite o documento da empresa: ")
		_, err := fmt.Scanln(&doc)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			continue // Tente novamente
		}

		if isValidDocument(doc) {
			break // Saia do loop se a entrada for válida
		} else {
			fmt.Println("Documento inválido. Tente novamente.")
		}
	}

	resultadoInstalador := salvarArquivo(zipURL, desktopPath, nomeInstalador, versao+".exe")

	//Executar o instalador do MicroSIP

	cmd := exec.Command(resultadoInstalador, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o instalador:%s ", err)
		return
	}

	// Executa o comando de Taskill
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

	fmt.Println("Agora vamos verificar se vc está na nossa base de dados, um momento")
	time.Sleep(1)
	// Salvando o arquivo ini na pasta \\AppData\\Roamming
	resultadoIni := salvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	// Editando o arquivo

	jsonfile, _ := database.BuscaPorDoc(doc, clientConfig)
	fmt.Println(jsonfile)
	if jsonfile != nil {
		fmt.Println("Encontrei você, ")
		fmt.Println("Os ramais que tenho vinculados a sua base são:")
		if err := json.Unmarshal(jsonfile, &clientConfig); err != nil {
			fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
			return
		}

	}

	//fmt.Println(&clientConfig)

	fmt.Println("Desculpe, não encontrei você . . . ")

	// Edição e Salvamento do arquivo .ini
	for _, config := range clientConfig {

		for i := range config.QuantRamais {
			fmt.Println(config.QuantRamais[i].Ramal)
		}

		fmt.Println("Porém os que não foram configurados ainda são:")

		for i := range config.QuantRamais {
			if config.QuantRamais[i].InUse == false {
				fmt.Println(config.QuantRamais[i].Ramal)
			}

		}

		var ramal string
		fmt.Print("Por favor informe agora, qual ramal você vai utilizar? ")
		_, err := fmt.Scanln(&ramal)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}
		//database.EditClient(config.ID, config)

		config.Ramal = ramal

		ramalInt, _ := strconv.Atoi(ramal)

		found := false

		for i := range config.QuantRamais {

			if config.QuantRamais[i].InUse != found {
				fmt.Println("Ramal sendo usado.")
				break
			}

			if config.QuantRamais[i].Ramal == ramalInt {
				config.QuantRamais[i].InUse = true
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Ramal não encontrado.")
			return
		}

		fmt.Println(config)

		// Codifique a estrutura de dados atualizada para JSON
		updatedJSON, err := json.Marshal(config.QuantRamais)
		if err != nil {
			fmt.Println("Erro ao serializar o JSON:", err)
			return
		}

		fmt.Printf(string(updatedJSON))

		database.EditClient(config.ID, config)
		//
		//
		//
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

}

func downloadFile(url string, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func salvarArquivo(link string, destination string, namePath string, extenssao string) string {

	// pegando o path C:\\%userprofile%
	//desktopPath, _ := os.UserHomeDir()

	// criando a pasta passando o path e o nome da pasta
	destinationFolder := filepath.Join(destination, namePath)
	if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
		fmt.Println("Erro ao criar a pasta na área de trabalho:", err)

	} // If da criação da pasta

	// salvando o arquivo
	arquivoTmp := filepath.Join(destinationFolder, namePath)
	if err := downloadFile(link, arquivoTmp+extenssao); err != nil {
		fmt.Println("Erro ao baixar o arquivo:", err)

	}
	fmt.Printf("Arquivo %s salvo com sucesso\n", namePath+extenssao)
	return arquivoTmp + extenssao

}

func isValidDocument(doc string) {

	//docInt, _ := strconv.Atoi(doc)
	var url string
	for i := range url {
		s := strconv.Itoa(i)
		url := "https://basesip.makesystem.com.br/clientes/" + s
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Erro ao fazer a solicitação HTTP:", err)
			return
		}
		defer response.Body.Close()
	}

	// Verifique o código de status da resposta HTTP
	///if response.StatusCode != http.StatusOK {
	//fmt.Println("Erro na resposta HTTP:", response.Status)
	//return
	//}

	//decoder := json.NewDecoder(response.Body)
	//if err := decoder.Decode(&docCliente); err != nil {
	//	fmt.Println("Erro ao decodificar JSON:", err)
	//	return
	//}

	// Itere sobre os clientes e imprima apenas o campo "doc"
	//for _, cliente := range docCliente {

	//	fmt.Println("Doc do cliente:", cliente.Doc)
	//}

}
