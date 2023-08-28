package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Jainish021/weather-app-go/routes"
	"github.com/Jainish021/weather-app-go/config"
)

func main() {
	r := mux.NewRouter()
	cfg := config.LoadConfig()

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	routes.WeatherAppRoutes(r)
	
	port := cfg.Port
	if port == "" {
        port = "3000"
    }

	fmt.Printf("Server is up on port %s.\n", port)

	log.Fatal(http.ListenAndServe("localhost:" + port, r))
}