package controller

import (
	"database/sql"
	"log"

	"github.com/nbanitama/web-application/core/model"
)

func GetBookController(dbConnection *sql.DB, pageId string) (model.Page, error) {
	log.Println("book controller")
	data := model.Page{}

	err := dbConnection.QueryRow("SELECT page_title,page_content,page_date "+
		"FROM pages WHERE id='?'", pageId).
		Scan(&data.Title, &data.Content, &data.Date)

	if err != nil {
		log.Println("Couldn't get page : +pageId")
		log.Println(err)
		return data, err
	}
	return data, nil
}
