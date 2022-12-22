package main

import (
	"famtree-api/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.SetupAuth()
	port := config.GetPort()
	router := config.GetRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		log.Fatal("Failed starting http server: ", err)
	}
}
