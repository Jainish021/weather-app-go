package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Jainish021/weather-app-go/routes"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	routes.WeatherAppRoutes(r)
	
	port := "3000"
	fmt.Printf("Server is up on port %s.\n", port)

	log.Fatal(http.ListenAndServe("localhost:3000", r))
}