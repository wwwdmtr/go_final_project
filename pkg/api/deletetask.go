package api

import (
	"errors"
	"net/http"

	"go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJsonError(w, errors.New("не указан идентификатор"))
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{})
}
