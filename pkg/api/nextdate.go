package api

import (
	"net/http"
	"time"

	"go_final_project/pkg/scheduler"
)

const DateFormat = scheduler.DateFormat

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	now, err := parseNow(r.FormValue("now"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	next, err := scheduler.NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(next))
}

func parseNow(value string) (time.Time, error) {
	if value == "" {
		return time.Now(), nil
	}

	return time.Parse(DateFormat, value)
}
