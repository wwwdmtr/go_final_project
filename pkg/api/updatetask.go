package api

import (
	"encoding/json"
	"errors"
	"go_final_project/pkg/db"
	"net/http"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJsonError(w, http.StatusBadRequest, err)
		return
	}

	if task.ID == "" {
		writeJsonError(w, http.StatusBadRequest, errors.New("не указан ID задачи"))
		return
	}

	if task.Title == "" {
		writeJsonError(w, http.StatusBadRequest, errors.New("не указан заголовок задачи"))
		return
	}

	if err := checkDate(&task); err != nil {
		writeJsonError(w, http.StatusBadRequest, err)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeInternalError(w, err)
		return
	}

	writeJson(w, http.StatusOK, map[string]any{})
}
