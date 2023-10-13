package routers

import (
	"github.com/gorilla/mux"
	"hello/controle"
	"hello/mid"
)

func HeandleRequest() {
	r := mux.NewRouter()
	r.Use(mid.ContentType)
	r.HandleFunc("/", controle.Home).Methods("Get")
	r.HandleFunc("/criar", controle.CriandoConta).Methods("Post")
}
