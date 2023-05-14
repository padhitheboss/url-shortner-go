package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/padhitheboss/url-shortner-go/pkg/model"
	"github.com/padhitheboss/url-shortner-go/pkg/routes"
)

var port = os.Getenv("PORT")

func main() {
	router := chi.NewRouter()

	defer model.DB.Close()
	model.M = make(map[string]model.Response)
	routes.RegisterRoute(router)
	fmt.Printf("Starting Server on Port %s", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}
