package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/harverone/go_final_project/pkg/api"
	"github.com/joho/godotenv"
)

var Defaultport string = "7450"
var EnvPortName string = "TODO_PORT"

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func SetPort(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Отсутсвует переменная с указанным портом, будет изпользован порт по умолчанию :%s\n", Defaultport)
		return fmt.Sprintf(":%s", Defaultport)
	}

	port := os.Getenv(name)

	if len(port) < 1 {
		log.Printf("Порт не указан, переменная пустая, будет изпользован порт по умолчанию :%s\n", Defaultport)
		return fmt.Sprintf(":%s", Defaultport)
	}
	if !IsInt(port) {
		log.Printf("Порт указан не корректно, будет использован порт по умолчанию :%s", Defaultport)
		return fmt.Sprintf(":%s", Defaultport)
	}

	log.Printf("Указан порт :%s\n", port)
	return fmt.Sprintf(":%s", port)
}

func StartServer() error {

	port := SetPort(EnvPortName)

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	log.Printf("Сервер запущен на порту %s\n", port)
	api.Init()
	return http.ListenAndServe(port, nil)
}
