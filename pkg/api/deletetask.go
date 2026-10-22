package api

import (
	"errors"
	"net/http"

	"go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJsonError(w, http.StatusBadRequest, errors.New("не указан идентификатор"))
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeInternalError(w, err)
		return
	}

	writeJson(w, http.StatusOK, map[string]any{})
}
