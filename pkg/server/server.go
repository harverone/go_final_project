package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/harverone/go_final_project/pkg/utils"
	"github.com/joho/godotenv"
)

// var Defaultport string = "7450"
// var EnvPort string = "TODO_PORT"

// //для проверки что в переменной окружения порт указан корректно(не дробь, буква и тд...)
// func IsInt(s string) bool {
// 	_, err := strconv.Atoi(s)
// 	return err == nil
// }

// load port from Env, use default if none or error
func SetPort(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Отсутсвует переменная с указанным портом, будет изпользован порт по умолчанию :%s\n", utils.Defaultport)
		return fmt.Sprintf(":%s", utils.Defaultport)
	}

	port := os.Getenv(name)

	if len(port) < 1 {
		log.Printf("Порт не указан, переменная пустая, будет изпользован порт по умолчанию :%s\n", utils.Defaultport)
		return fmt.Sprintf(":%s", utils.Defaultport)
	}
	if !utils.IsInt(port) {
		log.Printf("Порт указан не корректно, будет использован порт по умолчанию :%s", utils.Defaultport)
		return fmt.Sprintf(":%s", utils.Defaultport)
	}

	log.Printf("Указан порт :%s\n", port)
	return fmt.Sprintf(":%s", port)
}

func StartServer() {
	port := SetPort(utils.EnvPort)

	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	log.Printf("Сервер запущен на порту %s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
