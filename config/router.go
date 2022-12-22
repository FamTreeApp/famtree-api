package config

import (
	"famtree-api/controller"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Index).Methods("GET")
	router.HandleFunc("/families/", controller.GetFamilies).Methods("GET")
	router.HandleFunc("/check-id-family/{family-id}", controller.CheckIsFamilyIdAvailable).Methods("GET")
	router.HandleFunc("/auth/{provider}", controller.Login).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", controller.LoginCallback).Methods("GET")
	return router
}
