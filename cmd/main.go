package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/padhitheboss/url-shortner-go/pkg/model"
	"github.com/padhitheboss/url-shortner-go/pkg/routes"
)

var HOST = ""
var PORT = 3000

func main() {
	router := chi.NewRouter()
	model.M = make(map[string]model.Response)
	routes.RegisterRoute(router)
	fmt.Printf("Starting Server on Port %d", PORT)
	err := http.ListenAndServe(HOST+":"+strconv.Itoa(PORT), router)
	if err != nil {
		panic(err)
	}
}
