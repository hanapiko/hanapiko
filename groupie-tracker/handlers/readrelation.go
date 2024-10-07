package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
ReadRelations fetches relation data for a specific artist from the API.
It takes a baseURL and an id as parameters and returns a Relation struct.
The function sends a GET request to the specified URL and decodes the JSON response.
If successful and the relation is found, it returns the Relation.
Otherwise, it returns an error indicating either API issues or relation not found.
*/
func ReadRelations(baseURL, id string) (Relation, error) {
	res, err := http.Get(baseURL + id)
	if err != nil {
		return Relation{}, err
	}
	defer res.Body.Close()

	var data Relation
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return Relation{}, err
	}
	if data.ID == 0 {
		return Relation{}, fmt.Errorf("relation not found for ID: %s", id)
	}

	return data, nil
}
