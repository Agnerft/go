package model

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Produto struct {
	Id         string
	Nome       string
	Descricao  string
	Preco      string
	Quantidade string
}

func ReadFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err

	}

	return records, err
}

func ManipulationFile() []Produto {
	records, err := ReadFile("D:/teste.csv")

	if err != nil {
		log.Fatal(err)
	}

	produtos := []Produto{}

	for _, record := range records {

		p := Produto{
			Id:         record[0],
			Nome:       record[1],
			Descricao:  record[2],
			Preco:      record[3],
			Quantidade: record[4],
		}

		/*fmt.Printf("Id = %s Nome = %s Descrição = %s Preço = %s Quantidade = %s \n",
		p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
		*/

		produtos = append(produtos, p)
	}
	fmt.Print("\n")
	return produtos
}
