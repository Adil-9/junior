package main

import (
	// "fmt"

	"junior/internal/handlers"
	"junior/internal/logger"
	"net/http"
	"time"
	// "strconv"
)

func main() {
	time.Sleep(time.Second * 5) // when docker-compose waiting till postgres container runs and postgre is ready to accept connections (for ease purposes)
	logger.Init()               //creating logger
	logger.InfoLog.Println("Logger created")
	logger.InfoLog.Println("Starting applicatioin")

	handler := handlers.CreateHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleRequest)
	if err := http.ListenAndServe(":9000", mux); err != nil {
		logger.DebugLog.Fatal(err)
	} else {
		logger.InfoLog.Println("Server running, listening on http://localhost:9000/")
	}
}
