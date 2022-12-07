package main

import (
	"log"
	"net/http"

	"jobsearch/companies/axxonsoft/projects/3rdparty/internal/remote/routers"
)

func main() {
	TaskStorage := storage.NewTaskStorage()
	TaskService := service.NewTaskService(TaskStorage)
	TaskHandler := handlers.NewTaskHandler(TaskService)
	r := routers.NewRouter(TaskHandler)
	serverAddress := ":8080"
	log.Println("Serving on port", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, r))
}
