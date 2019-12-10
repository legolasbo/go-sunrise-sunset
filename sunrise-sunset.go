package sunriset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// APIData holds the data retrieved from the sunrise-sunset api.
type APIData struct {
	Sunrise                   string `json:"sunrise"`
	Sunset                    string `json:"sunset"`
	Solarnoon                 string `json:"solar_noon"`
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
func GetData(date, latitude, longitude string) (APIData, error) {
	apiURL := "https://api.sunrise-sunset.org/json?lat=" + latitude + "&lng=" + longitude + "&date=" + date + "&formatted=0"

	resp, err := http.Get(apiURL)
	if err != nil {
		return APIData{}, fmt.Errorf("Could not read from the api: %s", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIData{}, fmt.Errorf("Could not read response: %s", err)
	}

	d := results{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return APIData{}, fmt.Errorf("Could not decode json: %s", err)
	}

	return d.Data, nil
}

func parseTime(n, t string) (time.Time, error) {
	to, err := time.Parse("2006-01-02T15:04:05+00:00", t)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing "+n+": %s", err)
	}

	return to, nil
}

// GetSunrise parses the time string into a time struct.
func (d *APIData) GetSunrise() (time.Time, error) {
	return parseTime("Sunrise", d.Sunrise)
}

// GetSunset parses the time string into a time struct.
func (d *APIData) GetSunset() (time.Time, error) {
	return parseTime("Sunset", d.Sunset)
}

// GetSolarnoon parses the time string into a time struct.
func (d *APIData) GetSolarnoon() (time.Time, error) {
	return parseTime("Solarnoon", d.Solarnoon)
}

// GetCivilTwilightBegin parses the time string into a time struct.
func (d *APIData) GetCivilTwilightBegin() (time.Time, error) {
	return parseTime("CivilTwilightBegin", d.CivilTwilightBegin)
}

// GetCivilTwilightEnd parses the time string into a time struct.
func (d *APIData) GetCivilTwilightEnd() (time.Time, error) {
	return parseTime("CivilTwilightEnd", d.CivilTwilightEnd)
}

// GetNauticalTwilightBegin parses the time string into a time struct.
func (d *APIData) GetNauticalTwilightBegin() (time.Time, error) {
	return parseTime("NauticalTwilightBegin", d.NauticalTwilightBegin)
}

// GetNauticalTwilightEnd parses the time string into a time struct.
func (d *APIData) GetNauticalTwilightEnd() (time.Time, error) {
	return parseTime("NauticalTwilightEnd", d.NauticalTwilightEnd)
}

// GetAstronomicalTwilightBegin parses the time string into a time struct.
func (d *APIData) GetAstronomicalTwilightBegin() (time.Time, error) {
	return parseTime("AstronomicalTwilightBegin", d.AstronomicalTwilightBegin)
}

// GetAstronomicalTwilightEnd parses the time string into a time struct.
func (d *APIData) GetAstronomicalTwilightEnd() (time.Time, error) {
	return parseTime("AstronomicalTwilightEnd", d.AstronomicalTwilightEnd)
}
