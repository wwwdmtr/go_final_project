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
		writeJsonError(w, err)
		return
	}

	if task.ID == "" {
		writeJsonError(w, errors.New("не указан ID задачи"))
		return
	}

	if task.Title == "" {
		writeJsonError(w, errors.New("не указан заголовок задачи"))
		return
	}

	if err := checkDate(&task); err != nil {
		writeJsonError(w, err)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{})
}
