package database

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuscaPorDoc(doc string) ([]byte, error) {
	//doc := 12310400000182

	url := "http://localhost:3004/clientes?doc=" + doc
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
	fmt.Println(string(body))

	return body, nil
}