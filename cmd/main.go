package main

import (
	"jobsearch/companies/axxonsoft/projects/3rdparty/internal/remote/routers"
	"log"
	"net/http"
)

func main() {
	ReqStorage := storage.NewReqStorage()
	ReqService := service.NewReqService(ReqStorage)
	ReqHandler := handlers.NewReqHandler(ReqService)
	r := routers.NewRouter(ReqHandler)
	sa := ":8080"
	log.Println("Serving on port", sa)
	log.Fatal(http.ListenAndServe(sa, r))

}
