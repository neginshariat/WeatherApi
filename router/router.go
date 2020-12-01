package router

import (
	"openWapi/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/weather", middleware.WeatherUrl).Methods("GET", "OPTIONS")
	
	return router
}
