package main

import (
	"flag"
	"net/http"

	"api-nosql/config"
	"api-nosql/routes"
	"api-nosql/services"
)

var currManager services.Manager

func main() {
	flag.Parse()
	config.Settings()

	currManager = services.GetManager()
	r := routes.ApiRoutes()
	http.ListenAndServe(":"+config.Port, r)
}
