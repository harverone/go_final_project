package utils

import "strconv"

var Defaultport string = "7450"
var EnvPort string = "TODO_PORT"

// для проверки что в переменной окружения порт указан корректно
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
