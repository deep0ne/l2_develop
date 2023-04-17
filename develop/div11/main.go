package main

import (
	"div11/api"
	"div11/util"
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
	server.ListenAndServe()
}
