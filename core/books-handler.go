package core

import (
	"log"
	"net/http"
	"text/template"

	"github.com/nbanitama/web-application/core/controller"
)

func (m *Module) GetBookIndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hitting books index!!!..")
	param := r.URL.Query()

	log.Println("%+v", param)
	data, err := controller.GetBookController(m.dbConnection, "ji")
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't get page!")
	}

	t, _ := template.ParseFiles("file/templates/page.html")
	t.Execute(w, data)
}
