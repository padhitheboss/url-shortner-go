package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/padhitheboss/url-shortner-go/pkg/helper"
	"github.com/padhitheboss/url-shortner-go/pkg/model"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Executed")
	URLId := chi.URLParam(r, "Id")
	// URLid := helper.GetShortURLID(url)
	result, err := model.GetURL(URLId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Status{Error: err.Error()})
		return
	}
	url := helper.EnforceHTTP(result)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GenShortURL(w http.ResponseWriter, r *http.Request) {
	var body model.CreateRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&model.Status{Error: "invalid request body"})
		return
	}
	fmt.Println(body.URL)
	var shortUrl string
	if body.ShorternURL == "" {
		shortUrl = helper.GenShortURL()
		body.ShorternURL = shortUrl
	}
	var res model.Response
	res.URL = body.URL
	res.ShorternURL = body.ShorternURL
	res.UserId = "nil"
	res.Expiry = 30
	err = model.StoreURL(&body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Status{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&res)
}
