package main

import (
	"Loja/configs"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	err != nil {
	panic(err)
	}

	r := chi.NewRouter()

	
}


