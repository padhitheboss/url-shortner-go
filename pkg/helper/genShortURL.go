package helper

import (
	"os"
	"strings"

	"github.com/google/uuid"
)

func GenShortURL() string {
	// var id string
	// for _, ok := model.M[id]; !ok; {
	id := uuid.New().String()[:6]
	// }
	return id
}

// This checks if the domain that we are shorting
func RemoveDomainLoop(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	url = strings.Replace(url, "www.", "", 1)
	domain := strings.Split(url, "/")[0]
	return domain != os.Getenv("DOMAIN")
}

func GetShortURLID(url string) string {
	url = strings.Replace(url, "http://", "", 1)
	url = strings.Replace(url, "https://", "", 1)
	url = strings.Replace(url, "www.", "", 1)
	id := strings.Split(url, "/")[1]
	return id
}

func EnforceHTTP(url string) string {
	url = strings.Replace(url, "https://", "", 1)
	url = strings.Replace(url, "http://", "", 1)
	return "http://" + url
}
