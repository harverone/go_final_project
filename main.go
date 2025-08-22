package main

import (
	"log"
	// "net/http"

	"github.com/harverone/go_final_project/pkg/db"
	"github.com/harverone/go_final_project/pkg/server"
)

func main() {

	if err := db.Init(db.Dbpath); err != nil {
		log.Fatalf("Ошибка подключения базы данных: %v", err)
	}
	defer db.Close()

	if err := server.StartServer(); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
