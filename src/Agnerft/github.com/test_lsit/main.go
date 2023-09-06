package main

import (
	"net/http"

	"github.com/test_lsit/router"
)

func main() {
	//fmt.Print(model.ManipulationFile())

	router.Router()
	http.ListenAndServe(":8080", nil)
}

//	copy (SELECT * FROM produtos ORDER BY id DESC) to 'D:\teste.csv' with csv DELIMITER ';' HEADER
