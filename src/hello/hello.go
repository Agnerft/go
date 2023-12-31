package main

import (
	"fmt"
	"hello/controle"
	models "hello/models"
	"net/http"
	"os"
	"os/exec"
)

var clienteConfig []models.ClienteConfig
var clienteTeste []models.ClienteTeste

func main() {

	http.HandleFunc("/", controle.Home)
	// Criar uma rota para a página do formulário HTML.
	http.HandleFunc("/criar", controle.CriandoConta)
	//
	//

	// Iniciar o servidor na porta 8080.
	fmt.Println("Servidor está ouvindo na porta 8080...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
	//
	//
	//
	//
	//

	// COMEÇO DI INI
	//desktopPath, _ := os.UserHomeDir()

	//zipURL := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"
	//versao := "-3.21.3"
	// link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"
	//nomeInstalador := "MicroSIP"
	//link := "https://raw.githubusercontent.com/Agnerft/microsip/main/TESTE/MicroSIP1/MicroSIP.txt"

	//resultadoIni := utils.SalvarArquivo(link, desktopPath+"\\AppData\\Roaming\\", nomeInstalador, ".ini")

	// Pegando a busca por doc e setando num jsonfile
	//jsonfile, _ := database.BuscaPorDoc(44764891000128)

	//if err := json.Unmarshal(jsonfile, &clienteConfig); err != nil {
	//	fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
	//	return
	//}

	//fmt.Println(clienteConfig[len(clienteConfig)-1])

	url := "https://basesip.makesystem.com.br/clientes/"
	clienteExistente, err := models.ObterClienteExistente(1, url)
	if err != nil {
		fmt.Errorf("Erro para localizar cliente: %s", clienteExistente.Cliente)
	}

	ramais, err := clienteExistente.AdicionarRamais(7801, 5)
	if err != nil {
		fmt.Errorf("Não foi possível criar os ramais: %d", err)
	}

	err = models.AtualizarCliente(clienteExistente, fmt.Sprintf("%s/%d", url, clienteExistente.ID))
	if err != nil {
		fmt.Println("Erro ao atualizar cliente:", err)
		return
	}

	fmt.Println(ramais)

	//novoCliente := &models.ClienteConfig{}

	//novoCliente.AdicionarRamais(7800, 10)

	//fmt.Println(novoCliente.Cliente)

	// tests.AdicionarRamais(7801, 10)

	// models.CriarNovoCliente(tests, "http://localhost:3004/clientes")

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

func Comandos(resultadoInstalador string) {

	// EXECUTANDO O INSTALADOR
	cmd := exec.Command(resultadoInstalador, "/S")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Erro ao executar o instalador:%s ", err)
		return
	}

	// FINALIZAND O MICROSIP
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
}

/*
func criandoConta(w http.ResponseWriter, r *http.Request) {
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

	clienteCriado, _ := models.CriarNovoCliente(&configTest, "http://localhost:3004/clientes")

	fmt.Println(len(clienteCriado))

}

*/

func con(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "formulario.html")
}
