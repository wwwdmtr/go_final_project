package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go_final_project/pkg/db"
	"go_final_project/pkg/scheduler"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJsonError(w, err)
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

	id, err := db.AddTask(&task)
	if err != nil {
		writeJsonError(w, err)
		return
	}

	writeJson(w, map[string]any{"id": id})
}

func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(scheduler.DateFormat)
	}

	t, err := time.Parse(scheduler.DateFormat, task.Date)
	if err != nil {
		return err
	}

	if task.Repeat == "" {
		if afterNow(now, t) {
			task.Date = now.Format(scheduler.DateFormat)
		}
		return nil
	}

	next, err := scheduler.NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return err
	}

	if afterNow(now, t) {
		task.Date = next
	}

	return nil
}

// afterNow сравнивает только календарные даты, без учёта времени суток,
// чтобы дата "сегодня" не считалась "прошедшей" из-за текущего времени в now.
func afterNow(now, t time.Time) bool {
	y1, m1, d1 := now.Date()
	y2, m2, d2 := t.Date()

	nowDate := time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC)
	tDate := time.Date(y2, m2, d2, 0, 0, 0, 0, time.UTC)

	return nowDate.After(tDate)
}
