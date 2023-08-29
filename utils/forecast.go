package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/Jainish021/weather-app-go/config"
)

func Forecast(latitude, longitude string) (map[string]interface{}, error) {
	prodcfg := config.LoadProdConfig()
	ForecastAPIKey := prodcfg.ForecastAPIKey
	// if prodcfg.ForecastAPIKey == ""{
	// 	devcfg := config.LoadDevConfig()
	// 	ForecastAPIKey = devcfg.ForecastAPIKey
	// }

	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s,%s&units=f",
		ForecastAPIKey, url.QueryEscape(latitude), url.QueryEscape(longitude))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the weather service!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read response from weather service!")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("Unable to parse weather service response!")
	}

	if errMsg, ok := data["error"].(map[string]interface{}); ok {
		return nil, fmt.Errorf("Unable to find location! Reason: %s", errMsg["info"])
	}

	currentData := data["current"].(map[string]interface{})
	weatherData := map[string]interface{}{
		"temperature":       currentData["temperature"],
		"feelslike":         currentData["feelslike"],
		"humidity":          currentData["humidity"],
		"precip":            currentData["precip"],
		"weather_icon":      currentData["weather_icons"].([]interface{})[0],
		"weather_description": currentData["weather_descriptions"].([]interface{})[0],
		"wind_speed":        currentData["wind_speed"],
		"wind_dir":          currentData["wind_dir"],
		"pressure":          currentData["pressure"],
		"uv_index":          currentData["uv_index"],
		"visibility":        currentData["visibility"],
	}

	return weatherData, nil
}