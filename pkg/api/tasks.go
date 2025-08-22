package api

import (
	"net/http"

	"github.com/harverone/go_final_project/pkg/db"
)

const limit = 30

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(limit)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	if tasks == nil {
		writeJson(w, TasksResp{Tasks: []*db.Task{}}, http.StatusOK)
		return
	}
	writeJson(w, TasksResp{Tasks: tasks}, http.StatusOK)
}
