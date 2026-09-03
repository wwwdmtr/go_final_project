package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

type TaskResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	tasks, err := db.GetTasks(50, search)
	if err != nil {
		writeInternalError(w, err)
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJson(w, http.StatusOK, TaskResp{Tasks: tasks})

}
