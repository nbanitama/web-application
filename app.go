package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbanitama/web-application/core"
	"github.com/nbanitama/web-application/core/handler"
)

func main() {
	log.Println("Starting the application")

	module, err := core.InitializeModule()
	if err != nil {
		log.Println("Error happened when initializing module")
		log.Println(err)
	} else {
		fmt.Printf("%+v", module)
		http.HandleFunc("/index", handler.ServeIndexFileHandler)
		http.HandleFunc("/books", handler.GetBookIndexHandler)

		http.ListenAndServe(":9999", nil)
	}
}
