package main

import (
	"net/http"

	"github.com/jacolpn/go-web-products/routes"
)

func main() {
	routes.CarregaRotas()

	http.ListenAndServe(":8000", nil)
}
