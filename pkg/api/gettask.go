package api

import (
	"go_final_project/pkg/db"
	"net/http"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, err := db.GetTask(id)
	if err != nil {
		writeJsonError(w, err)
		return
	}

	writeJson(w, task)
}
