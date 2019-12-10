package sunriset

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// APIData holds the data retrieved from the sunrise-sunset api.
type APIData struct {
	Sunrise                   string `json:"sunrise"`
	Sunset                    string `json:"sunset"`
	SolarNoon                 string `json:"solar_noon"`
	DayLength                 string `json:"day_length"`
	CivilTwilightBegin        string `json:"civil_twilight_begin"`
	CivilTwilightEnd          string `json:"civil_twilight_end"`
	NauticalTwilightBegin     string `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       string `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
}

type results struct {
	Data   APIData `json:"results"`
	Status string  `json:"status"`
}

// GetData retrieves the data from the sunrise-sunset api.
func GetData(date, latitude, longitude string) APIData {
	apiURL := "https://api.sunrise-sunset.org/json?lat=" + latitude + "&lng=" + longitude + "&date=" + date + "&formatted=0"

	resp, _ := http.Get(apiURL)
	body, _ := ioutil.ReadAll(resp.Body)

	d := results{}
	json.Unmarshal(body, &d)

	return d.Data
}
