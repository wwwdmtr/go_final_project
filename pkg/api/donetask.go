package api

import (
	"errors"
	"net/http"
	"time"

	"go_final_project/pkg/db"
	"go_final_project/pkg/scheduler"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJsonError(w, errors.New("не указан идентификатор"))
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJsonError(w, err)
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJsonError(w, err)
			return
		}
		writeJson(w, map[string]any{})
		return
	}

	nextDate, err := scheduler.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJsonError(w, err)
		return
	}

	task.Date = nextDate
	if err := db.UpdateTask(task); err != nil {
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{})
}
