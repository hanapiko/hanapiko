package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
ReadArtist fetches artist information from the API.
It takes a baseUrl and an id as parameters and returns an Artist struct.
The function sends a GET request to the specified URL and decodes the JSON response.
If successful and the artist is found, it returns the Artist.
Otherwise, it returns an error indicating either API issues or artist not found.
*/
func ReadArtist(baseUrl, id string) (Artist, error) {
	baseUrl = fmt.Sprintf("%s%s", baseUrl, id)
	response, err := http.Get(baseUrl)
	if err != nil {
		return Artist{}, err
	}
	defer response.Body.Close()

	var artist Artist
	if response.StatusCode == http.StatusOK {
		err = json.NewDecoder(response.Body).Decode(&artist)
		if err != nil {
			return Artist{}, err
		}
		if artist.ID == 0 {
			return Artist{}, fmt.Errorf("Artist not found")
		}
	} else {
		return Artist{}, fmt.Errorf("API returned status code: %d", response.StatusCode)
	}
	return artist, nil
}
