package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/padhitheboss/url-shortner-go/pkg/helper"
	"github.com/padhitheboss/url-shortner-go/pkg/model"
)

func GetURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Executed")
	URLId := chi.URLParam(r, "Id")
	// URLid := helper.GetShortURLID(url)
	result, ok := model.M[URLId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&model.Status{Error: "url not found"})
		return
	}
	url := helper.EnforceHTTP(result.URL)
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
	shortUrl := helper.GenShortURL()
	var res model.Response
	res.URL = body.URL
	res.ShorternURL = shortUrl
	res.UserId = "nil"
	res.Expiry = 30 * time.Minute
	model.M[shortUrl] = res
	json.NewEncoder(w).Encode(&res)
}
