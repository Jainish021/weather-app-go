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

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	routes.WeatherAppRoutes(r)
	
	prodcfg := config.LoadProdConfig()
	port := prodcfg.Port
	// if prodcfg.Port == ""{
	// 	devcfg := config.LoadDevConfig()
	// 	port = devcfg.Port
	// }

	fmt.Printf("Server is up on port %s.\n", port)

	log.Fatal(http.ListenAndServe(":" + port, r))
}