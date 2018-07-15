package core

import (
	"log"
)

type Module struct {
	num int
}

func InitializeModule() (*Module, error) {
	log.Println("Initializing the module...")
	module := Module{
		num: 1,
	}

	return &module, nil
}
