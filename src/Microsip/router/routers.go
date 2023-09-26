package router

import (
	"Microsip/controle"
	"Microsip/mid"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HeandleRequest() {
	r := mux.NewRouter()
	r.Use(mid.ContentType)
	r.HandleFunc("/", controle.Home).Methods("GET")
	r.HandleFunc("/execute", controle.SaveExe).Methods("GET")
	r.HandleFunc("/{doc}", controle.FindUser).Methods("GET")

	fmt.Println("Servidor está ouvindo na porta 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
