package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Init() {

	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneHandler)
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addTaskHandler(w, r)
	case "GET":
		getTaskHandler(w, r)
	case "PUT":
		updateTaskHandler(w, r)
	case "DELETE":
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "Неизвестный метод", http.StatusMethodNotAllowed)
	}
}

func writeJson(w http.ResponseWriter, data any, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	dataJson, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка сериализации: %v", err), http.StatusInternalServerError)
	}
	w.Write(dataJson)
}
