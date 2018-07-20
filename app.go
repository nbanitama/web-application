package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nbanitama/web-application/core"
)

func main() {
	log.Println("Starting the application")

	module, err := core.InitializeModule()

	if err != nil {
		log.Println("Error happened when initializing module")
		log.Println(err)
	} else {
		fmt.Printf("%+v", module)
		http.HandleFunc("/index", module.ServeIndexFileHandler)
		http.HandleFunc("/page", module.GetBookIndexHandler)

		http.ListenAndServe(":9999", nil)
	}
}
