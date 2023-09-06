package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/ler_arquivo/user"
)

func main() {

	leres, err := lerArquivo("desafio_junior.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, ler := range leres {

		usuario := user.Usuario{
			Nome:           ler[0],
			Telefone:       ler[1],
			Documento:      ler[2],
			DataNascimento: ler[3],
		}

		fmt.Println(usuario)
	}

}

func lerArquivo(arquivo string) ([][]string, error) {

	f, err := os.Open("desafio_junior.csv")

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	ler, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return ler, nil

}
