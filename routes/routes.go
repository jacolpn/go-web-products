package routes

import (
	"net/http"

	"github.com/jacolpn/go-web-products/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Editar)
	http.HandleFunc("/update", controllers.Update)
}
