package main

import (
	"div11/api"
	"div11/util"
	"log"
	"net/http"
)

func main() {
	config := util.NewConfig(":8080")
	api := api.NewServer(config)
	router := api.NewRouter()
	server := http.Server{
		Addr:    api.Config.ServerAddress,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
