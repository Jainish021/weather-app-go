package config

import(
	"os"
    "github.com/Jainish021/weather-app-go/models"
)

func LoadProdConfig() models.ProdConfig {
    return models.ProdConfig{
        Port: os.Getenv("PORT"),
        ForecastAPIKey: os.Getenv("ForecastAPIKey"),
		GeocodeAPIKey: os.Getenv("GeocodeAPIKey"),
    }
}