package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"
	"html/template"
	"github.com/Jainish021/weather-app-go/utils"
	"github.com/Jainish021/weather-app-go/models"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title: "Weather",
		Name :  "Jainish Adesara",
	}

	tmpl, _ := template.ParseFiles("templates/index.html", "templates/partials/head.html", "templates/partials/header.html", "templates/partials/footer.html")
	tmpl.Execute(w, data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title: "About",
		Name:  "Jainish Adesara",
	}

	tmpl, _ := template.ParseFiles("templates/about.html", "templates/partials/head.html", "templates/partials/header.html", "templates/partials/footer.html")
	tmpl.Execute(w, data)
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		HelpText: "This is a help page.",
		Title:    "Help",
		Name:     "Jainish Adesara",
	}

	tmpl, _ := template.ParseFiles("templates/help.html", "templates/partials/head.html", "templates/partials/header.html", "templates/partials/footer.html")
	tmpl.Execute(w, data)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	address := r.URL.Query().Get("address")
	
	if address == "" {
		http.Error(w, "Address must be provided.", http.StatusBadRequest)
		return
	}

	locationData, err := utils.Geocode(address)
	if err != nil {
		data := models.PageData{
			ErrorMessage: "Something went wrong. Please try again.",
		}
		
		res, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

		return
	}

	weatherData, err := utils.Forecast(fmt.Sprintf("%f", locationData["latitude"]), fmt.Sprintf("%f", locationData["longitude"]))
	if err != nil {
		data := models.PageData{
			ErrorMessage: "Something went wrong. Please try again.",
		}
		
		res, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		
		return
	}

	data := models.WeatherData{
		Location: locationData["location"].(string),
		Temperature: weatherData["temperature"].(float64),
		Feelslike: weatherData["feelslike"].(float64),
		Humidity: weatherData["humidity"].(float64),
		Precip: weatherData["precip"].(float64),
		WeatherIcon: weatherData["weather_icon"].(string),
		WeatherDescription: weatherData["weather_description"].(string),
		WindSpeed: weatherData["wind_speed"].(float64),
		WindDir: weatherData["wind_dir"].(string),
		Pressure: weatherData["pressure"].(float64),
		UVIndex: weatherData["uv_index"].(float64),
		Visibility: weatherData["visibility"].(float64),
	}

	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:       "404",
		ErrorMessage: "Page not found",
		Name:        "Jainish Adesara",
	}

	tmpl, _ := template.ParseFiles("templates/404.html", "templates/partials/head.html", "templates/partials/header.html", "templates/partials/footer.html")
	tmpl.Execute(w, data)
}