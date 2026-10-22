package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.Println("writeJson: marshal error:", err)
		http.Error(w, `{"error":"внутренняя ошибка"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(resp)
}

func writeJsonError(w http.ResponseWriter, status int, err error) {
	writeJson(w, status, map[string]any{"error": err.Error()})
}

// writeInternalError логирует настоящую ошибку (может содержать чувствительные
// детали — запросы к БД и т.п.) и отдаёт клиенту только заглушку.
func writeInternalError(w http.ResponseWriter, err error) {
	if errors.Is(err, sql.ErrNoRows) {
		writeJsonError(w, http.StatusNotFound, errors.New("задача не найдена"))
		return
	}

	log.Println("internal error:", err)
	writeJsonError(w, http.StatusInternalServerError, errors.New("внутренняя ошибка"))
}
