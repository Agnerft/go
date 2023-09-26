package controle

import (
	"Microsip/database"
	"Microsip/models"
	"Microsip/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"os"

	"net/http"
)

var (
	versao = "-3.21.3"
	// URL do arquivo ZIP
	zipURL         = "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	link           = "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	nomeInstalador = "MicroSIP"
	clientConfig   []models.ClienteConfig
	desktopPath, _ = os.UserHomeDir()
)

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "screen.html")
}

func SaveExe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)

		return
	}

	resultadoInstalador := utils.SalvarArquivo(zipURL, desktopPath, nomeInstalador, versao+".exe")

	//Executar o instalador do MicroSIP

	utils.Comandos(resultadoInstalador)
}

func SaveIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}

	resultadoIni := utils.SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	for _, config := range clientConfig {
		linkCompleto := config.GrupoRecurso + config.LinkGvc + config.Porta
		utils.Editor(resultadoIni, 2, "label="+config.Ramal)
		utils.Editor(resultadoIni, 3, "server="+linkCompleto)
		utils.Editor(resultadoIni, 4, "proxy="+linkCompleto)
		utils.Editor(resultadoIni, 5, "domain="+linkCompleto)
		utils.Editor(resultadoIni, 6, "username="+config.Ramal)
		utils.Editor(resultadoIni, 7, "password="+config.Ramal+config.Senha)
		utils.Editor(resultadoIni, 8, "authID="+config.Ramal)
	}

	fmt.Println("Cliente salvo e ajustado")

}

func CriandoConta(w http.ResponseWriter, r *http.Request) {
	// Verificar se a solicitação é do tipo POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar os dados JSON do corpo da solicitação.
	//var configTest models.ClienteTeste
	configTest := models.NewClienteTeste()
	err := json.NewDecoder(r.Body).Decode(&configTest)
	if err != nil {
		http.Error(w, "Erro ao decodificar os dados JSON", http.StatusBadRequest)
		return
	}

	// Aqui, você pode processar os dados do usuário, por exemplo, salvá-los em um banco de dados.
	// Vamos apenas exibir os dados recebidos para este exemplo.
	fmt.Printf("Novo usuário criado: %+v\n", configTest)

	// Responder com uma mensagem de sucesso.
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuário criado com sucesso!\n")

	clienteCriado, err := models.CriarNovoCliente(&configTest, "https://basesip.makesystem.com.br/clientes")
	if err != nil {
		fmt.Println("Erro")
	}

	fmt.Println(len(clienteCriado))

	//http.ServeFile(w, r, "hello.html")

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	// Verificar se a solicitação é do tipo GET.

	vars := mux.Vars(r)

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	jsonfile, _ := database.BuscaPorDoc(vars["doc"], clientConfig)
	fmt.Println(jsonfile)

	if jsonfile == nil {
		fmt.Println("Desculpe, não encontrei você . . . ")
		return
	}

	htmlResponse := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Minha Página HTML</title>
        </head>
        <body>
            <h1>Olá, mundo!</h1>
            <p>Esta é uma resposta HTML de exemplo.</p>
        </body>
        </html>
    `

	fmt.Fprintf(w, htmlResponse)

}
