package main

import (
	"fmt"
	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	port := "7540"
	envPort := os.Getenv("TODO_PORT")
	if envPort != "" {
		port = envPort
	}

	dbFile := "scheduler.db"
	envDBFile := os.Getenv("TODO_DBFILE")
	if envDBFile != "" {
		dbFile = envDBFile
	}

	err := db.InitDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	api.Init()
	fmt.Printf("Server is running on port %s\n", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
