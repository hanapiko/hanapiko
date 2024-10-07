package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
ReadDate fetches date information from the API.
It takes a baseURL and an id as parameters and returns a DateEntry struct.
The function sends a GET request to the specified URL and decodes the JSON response.
If successful, it returns the DateEntry. Otherwise, it returns an error.
*/
func ReadDate(baseURL, id string) (DateEntry, error) {
	url := fmt.Sprintf("%s%s", baseURL, id)
	response, err := http.Get(url)
	if err != nil {
		return DateEntry{}, err
	}
	defer response.Body.Close()

	var artist DateEntry
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artist)
		if err != nil {
			return DateEntry{}, err
		}
	} else {
		return DateEntry{}, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artist, nil
}
