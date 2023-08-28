package main

import (
	"os"
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
	
	port := os.Getenv("PORT")
	if port == "" {
        port = cfg.Port
    }

	fmt.Printf("Server is up on port %s.\n", port)

	log.Fatal(http.ListenAndServe("localhost:" + port, r))
}