package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
ReadLocation fetches location information from the API.
It takes a baseURL and an id as parameters and returns a Location struct.
The function sends a GET request to the specified URL and decodes the JSON response.
If successful, it returns the Location. Otherwise, it returns an error.
*/
func ReadLocation(baseURL, id string) (Location, error) {
	baseURL = fmt.Sprintf("%s%s", baseURL, id)
	response, err := http.Get(baseURL)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()

	var artist Location
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artist)
		if err != nil {
			return Location{}, err
		}
	} else {
		return Location{}, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artist, nil
}
