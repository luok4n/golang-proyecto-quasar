package routes

import (
	"MeLi/controllers"

	"github.com/gorilla/mux"
)

func SetServicesRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/topsecret", controllers.PostTopSecret).Methods("POST")
	subRoute.HandleFunc("/topsecret/{satellite_name}", controllers.TopSecret)
}
