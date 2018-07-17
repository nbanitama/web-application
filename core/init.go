package core

import (
	"database/sql"
	"log"

	"github.com/nbanitama/web-application/tool"
)

type Module struct {
	num          int
	dbConnection *sql.DB
}

func InitializeModule() (*Module, error) {
	log.Println("Initializing the module...")
	module := Module{
		num: 1,
	}

	db, err := tool.InitializeDB()
	if err != nil {
		log.Println("Couldn't connect!!...")
		return nil, err
	}

	log.Println("DB connected...")
	module.dbConnection = db
	return &module, nil
}
