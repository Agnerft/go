package utils

import (
	"encoding/json"
	"fmt"
	"hello/database"
	"hello/models"
	"os"
	"path/filepath"
)

var clienteConfig []models.ClienteConfig
var clienteTeste []models.ClienteTeste

func SaveFale() {

	//COMEÇO DI INI
	desktopPath, _ := os.UserHomeDir()

	zipURL := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	versao := "-3.21.3"
	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	nomeInstalador := "MicroSIP"
	link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"

	resultadoIni := SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	//Pegando a busca por doc e setando num jsonfile
	jsonfile, _ := database.BuscaPorDoc(44764891000128)

	if err := json.Unmarshal(jsonfile, &clienteConfig); err != nil {
		fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
		return
	}
}

func SalvarArquivo(link string, destination string, namePath string, extenssao string) string {

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
