package handler

import (
	"log"
	"net/http"
)

func GetBookIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hitting books index!!!..")
}
