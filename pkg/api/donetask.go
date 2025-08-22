package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/harverone/go_final_project/pkg/db"
)

func doneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка получения задачи: %v", err)}, http.StatusInternalServerError)
		return
	}
	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка удаления задачи: %v", err)}, http.StatusInternalServerError)
			return
		}
		writeJson(w, map[string]int64{}, http.StatusOK)
	} else {
		nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка вычисления следующей даты: %v", err)}, http.StatusInternalServerError)
			return
		}
		err = db.UpdateTaskDate(id, nextDate)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("ошибка обновления даты: %v", err)}, http.StatusInternalServerError)
			return
		}
		writeJson(w, map[string]int64{}, http.StatusOK)
	}
}
