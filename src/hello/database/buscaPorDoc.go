package database

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func BuscaPorDoc(doc int) ([]byte, error) {
	//doc := 12310400000182
	// ajustar porta
	url := "https://basesip.makesystem.com.br/clientes?doc=" + strconv.Itoa(doc)
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("Erro ao finalizar a abertura da requisição: %d", res.StatusCode)
		}

	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return

	}
	//fmt.Println(string(body))

	return body, nil
}
