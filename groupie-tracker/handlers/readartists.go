package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
ReadArtists fetches a list of artists from the API.
It takes a url as a parameter and returns a slice of Artist structs.
The function sends a GET request to the specified URL and decodes the JSON response.
If successful, it returns the slice of Artists. Otherwise, it returns an error.
*/
func ReadArtists(url string) ([]Artist, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []Artist
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artists)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artists, nil
}
