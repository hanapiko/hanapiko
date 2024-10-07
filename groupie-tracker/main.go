package main

import (
	"net/http"
	"os"

	api "groupie/handlers"
)

func main() {
	if len(os.Args) != 1 {
		return
	}
	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/locations/", api.LocationHandler)
	http.HandleFunc("/artists/", api.ArtistsHandler)
	http.HandleFunc("/artist/", api.ArtistHandler)
	http.HandleFunc("/relation/", api.RelationHandler)
	http.HandleFunc("/dates/", api.DateHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
