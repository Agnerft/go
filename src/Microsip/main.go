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
	"time"
)

var clientConfig []models.ClienteConfig

func main() {

	versao := "-3.21.3"
	// URL do arquivo ZIP
	zipURL := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	nomeInstalador := "MicroSIP"

	desktopPath, _ := os.UserHomeDir()

	var doc string

	fmt.Print("Digite o documento da empresa: ")
	_, err := fmt.Scanln(&doc)

	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
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

	err = cmd3.Run()
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

	jsonfile, _, nomeCliente := database.BuscaPorDoc(doc, clientConfig)
	fmt.Println("Encontrei, " + nomeCliente)
	if err := json.Unmarshal(jsonfile, &clientConfig); err != nil {
		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
		return
	}

	// Edição e Salvamento do arquivo .ini
	for _, config := range clientConfig {
		var ramal string
		fmt.Print("Qual ramal você vai utilizar? ")
		_, err := fmt.Scanln(&ramal)
		if err != nil {
			fmt.Println("Erro ao ler a entrada:", err)
			return
		}

		config.Ramal = ramal
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
