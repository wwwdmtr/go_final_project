package api

import "net/http"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", authMiddleware(taskHandler))
	http.HandleFunc("/api/tasks", authMiddleware(getTasksHandler))
	http.HandleFunc("/api/task/done", authMiddleware(doneTaskHandler))
	http.HandleFunc("/api/signin", signinHandler)
}
