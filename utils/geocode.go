package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Geocode(address string) (map[string]interface{}, error) {
	accessToken := "pk.eyJ1IjoiamFpbmlzaDAyMSIsImEiOiJja3BvZGV5aHg0bmNvMm9ueHM0NGwyZ2p4In0.yKwgCnboa4q2d1Lss3BmXg"
	geocodeURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s&limit=1",
		url.QueryEscape(address), accessToken)

	resp, err := http.Get(geocodeURL)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to location service!")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read response from location service!")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("Unable to parse location service response!")
	}

	features := data["features"].([]interface{})
	if len(features) == 0 {
		return nil, fmt.Errorf("Unable to find location!")
	}

	firstFeature := features[0].(map[string]interface{})
	latitude := firstFeature["center"].([]interface{})[1].(float64)
	longitude := firstFeature["center"].([]interface{})[0].(float64)
	placeName := firstFeature["place_name"].(string)

	locationData := map[string]interface{}{
		"latitude":  latitude,
		"longitude": longitude,
		"location":  placeName,
	}

	return locationData, nil
}