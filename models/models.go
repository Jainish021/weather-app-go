package models

type DevConfig struct {
    Port     string
    ForecastAPIKey   string
	GeocodeAPIKey	string
}

type ProdConfig struct {
    Port     string
    ForecastAPIKey   string
	GeocodeAPIKey	string
}

type PageData struct{
	Title string
	Name string
	HelpText string
	ErrorMessage string
}

type WeatherData struct{
	Location string
	Temperature float64
	Feelslike float64
	Humidity float64
	Precip float64
	WeatherIcon string
	WeatherDescription string
	WindSpeed float64
	WindDir string
	Pressure float64
	UVIndex float64
	Visibility float64
}