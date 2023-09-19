package database

import (
	"encoding/json"
	"fmt"
	"github.com.br/agnerft/models"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	login string
	senha string
)

func BuscaPorDoc(doc string, c []models.ClienteConfig) ([]byte, error, string) {
	//doc := 12310400000182
	login := "root"
	senha := "agner102030"
	url := "https://" + login + ":" + senha + "@" + "basesip.makesystem.com.br/clientes?doc=" + doc
	method := "GET"

	//fmt.Println(url)

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return

	}
	//fmt.Println(string(body))

	respBody := string(body)
	//teste := strings.Map(strings.Split(respBody, ","))
	//nomeClient :=

	contagem := strings.Count(respBody, "doc")
	if contagem > 1 {
		fmt.Println("Existe clientes repetidos.")

		return nil, err, ""
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Printf("Não deu para fazer a junção. \n")

		return nil, err, ""
	}

	fmt.Println(c[0].Cliente)

	return body, nil, c[0].Cliente
}
