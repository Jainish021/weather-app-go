package config

import(
	"os"
)

type ProdConfig struct {
    Port     string
    ForecastAPIKey   string
	GeocodeAPIKey	string
}

func LoadProdConfig() ProdConfig {
    return ProdConfig{
        Port: os.Getenv("PORT"),
        ForecastAPIKey: os.Getenv("ForecastAPIKey"),
		GeocodeAPIKey: os.Getenv("GeocodeAPIKey"),
    }
}