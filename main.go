package main

import (
	// "log"
	// "net/http"

	"github.com/harverone/go_final_project/pkg/server"
)

func main() {

	// port := server.SetPort(server.EnvPort)

	// fs := http.FileServer(http.Dir("./web"))
	// http.Handle("/", fs)

	// log.Printf("Сервер запущен на порту %s\n", port)

	// log.Fatal(http.ListenAndServe(port, nil))
	server.StartServer()
}
