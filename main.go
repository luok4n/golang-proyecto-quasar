package main

import (
	controllers "MeLi/controllers"
	"MeLi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	controllers.SetCoordinates()

	router := mux.NewRouter()
	routes.SetServicesRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	server.ListenAndServe()

}
