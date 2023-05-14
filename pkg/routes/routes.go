package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/padhitheboss/url-shortner-go/pkg/controller"
)

func RegisterRoute(r chi.Router) {
	r.Get("/{Id}", controller.GetURL)
	r.Post("/generate", controller.GenShortURL)
	fmt.Println("Routing Executed")
}
