package routes

import(
	"github.com/gorilla/mux"
	"github.com/Jainish021/weather-app-go/controllers"
)

var WeatherAppRoutes = func(router *mux.Router){
	router.HandleFunc("/", controllers.IndexHandler).Methods("GET")
	router.HandleFunc("/about", controllers.AboutHandler).Methods("GET")
	router.HandleFunc("/help", controllers.HelpHandler).Methods("GET")
	router.HandleFunc("/weather", controllers.WeatherHandler).Methods("GET")
	router.PathPrefix("/").HandlerFunc(controllers.NotFoundHandler)
}