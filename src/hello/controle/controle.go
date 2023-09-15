package controle

import (
	"encoding/json"
	"fmt"
	"hello/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "screen.html")
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

	clienteCriado, err := models.CriarNovoCliente(&configTest, "http://localhost:3004/clientes")
	if err != nil {
		fmt.Println("Erro")
	}

	fmt.Println(len(clienteCriado))

	//http.ServeFile(w, r, "hello.html")

}
