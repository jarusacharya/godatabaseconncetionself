package routers

import (
	controller "databaseserverintegration/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/course/{id}", controller.GetACourse).Methods("GET")
	router.HandleFunc("/api/course", controller.GetAllCourse).Methods("GET")
	router.HandleFunc("/api/course", controller.InsertACourse).Methods("POST")
	router.HandleFunc("/api/course/{id}", controller.UpdateACourse).Methods("PUT", "PATCH")
	router.HandleFunc("/api/course/{id}", controller.DeleteACourse).Methods("DELETE")
	router.HandleFunc("/api/course", controller.DeleteAllCourse).Methods("DELETE")
	return router
}
