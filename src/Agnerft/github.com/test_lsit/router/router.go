package router

import (
	"net/http"

	"github.com/test_lsit/controller"
)

func Router() {

	http.HandleFunc("/", controller.ListProducts)
}
