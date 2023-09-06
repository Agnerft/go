package handlers

import (
	"Loja/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	id, err := models.Insert(todo)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Não foi possível fazer o Dedode. Erro %v", err),
		}

	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID:%d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

// N TENDO ERRO ELE PEGA O id, err := chamando o
//models.Insert(E passando o que eu quero inserir)
// crio um var resp faço um map[string]any
/*Faço um if err != nil {
	resp = ma[string]any {
		"Error": true,
		"Message": MEnsagem que deu erro
	} else
}

dados da jrsolution
microsip: jrsolution.gvctelecom.com.br:5131
link da telefonia: jrsolution.gvctelecom.com.br:1163

*/
