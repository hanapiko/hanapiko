package api

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

var errorTemplate *template.Template

/*
Init initializes the error template for the application.
It attempts to parse the error.html template file. If parsing fails,
it creates a simple fallback template to ensure error rendering.
This function should be called once at the start of the application.
*/
func Init() {
	var err error
	errorTemplate, err = template.ParseFiles("template/error.html")
	if err != nil {
		// log.Printf("Warning: Error parsing error template: %v", err)
		// Create a simple fallback template
		errorTemplate = template.Must(template.New("error").Parse(`
            <html><body>
            <h1>Error {{.Code}}</h1>
            <p>{{.Message}}</p>
            </body></html>
        `))
		//  log.Println("Error parsing, using fallback template")
	}
}

/*
renderError handles the rendering of error pages.
It sets the HTTP status code, executes the error template with the provided status and message,
and logs any errors that occur during template execution.
Parameters:
  - w: http.ResponseWriter to write the response
  - status: HTTP status code for the error
  - message: Error message to display
*/
func renderError(w http.ResponseWriter, status int, message string) {
	Init()
	w.WriteHeader(status)
	err := errorTemplate.Execute(w, struct {
		Code    int
		Message string
	}{
		Code:    status,
		Message: message,
	})
	if err != nil {
		log.Printf("Error rendering error template: %v", err)
	}
}

/*
HomeHandler manages requests to the home page of the application.
It checks if the requested path is the root ("/") and if the HTTP method is GET.
If these conditions are not met, it renders appropriate error pages.
Otherwise, it parses and executes the home.html template.

Parameters:
  - w: http.ResponseWriter to write the response
  - r: *http.Request containing the request details
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	// Parse the homepage template
	temp, err := template.ParseFiles("template/home.html") // Ensure you have home.html in the template directory
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	// Execute the template and write to the response
	err = temp.Execute(w, nil) // No data is passed to the homepage template
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

/*
ArtistsHandler manages requests to the artists listing page.
It verifies the correct URL path and HTTP method, then fetches and displays
the list of artists. If any errors occur during this process, it renders
appropriate error pages.

Parameters:
  - w: http.ResponseWriter to write the response
  - r: *http.Request containing the request details
*/
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists/" {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	templatePath := filepath.Join("template", "artists.html")
	temp1, err := template.ParseFiles(templatePath)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	result, err := ReadArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error fetching artists")
		return
	}

	err = temp1.Execute(w, result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

/*
ArtistHandler manages requests for individual artist pages.
It checks for the correct HTTP method and URL format, extracts the artist ID
from the URL, fetches the artist data, and renders it using the artist template.
If any errors occur during this process, it renders appropriate error pages.

Parameters:
  - w: http.ResponseWriter to write the response
  - r: *http.Request containing the request details
*/
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/artist/") || len(strings.Split(r.URL.Path, "/")) != 3 {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "Artist ID not found")
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/artist.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}
	baseURL := "https://groupietrackers.herokuapp.com/api/artists/"

	result, err := ReadArtist(baseURL, id)
	if err != nil {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	if result.ID == 0 {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	err = temp1.Execute(w, result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

/*
LocationHandler manages requests for location information of artists.
It verifies the HTTP method, extracts the location ID from the URL,
fetches the location data, and renders it using the locations template.
If any errors occur during this process, it renders appropriate error pages.

Parameters:
  - w: http.ResponseWriter to write the response
  - r: *http.Request containing the request details
*/
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "Artist ID not found")
		return
	}
	id := id1[len(id1)-1]

	temp1, err := template.ParseFiles("template/locations.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/locations/"

	Result, err := ReadLocation(url, id)
	if err != nil {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	if Result.ID == 0 {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	err = temp1.Execute(w, Result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

/*
DateHandler manages requests for concert date information of artists.
It verifies the HTTP method, extracts the artist ID from the URL,
fetches the date data, and renders it using the dates template.
If any errors occur during this process, it renders appropriate error pages.

Parameters:
  - w: http.ResponseWriter to write the response
  - r: *http.Request containing the request details
*/
func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Wrong method")
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "dates" {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	id := parts[2]
	if id == "" {
		renderError(w, http.StatusBadRequest, "Artist ID not found")
		return
	}

	temp1, err := template.ParseFiles("template/dates.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template")
		return
	}

	baseURL := "https://groupietrackers.herokuapp.com/api/dates/"
	Result, err := ReadDate(baseURL, id)
	if err != nil {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	if Result.ID == 0 {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	err = temp1.Execute(w, Result)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error executing template")
	}
}

/*
RelationHandler manages requests for relation information of artists.
It verifies the HTTP method, extracts the relation ID from the URL,
fetches the relation data, and renders it using the relation template.
If any errors occur during this process, it renders appropriate error pages.

Parameters:
  - w: http.ResponseWriter to write the response
  - r: *http.Request containing the request details
*/
func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/relation/") || len(strings.Split(r.URL.Path, "/")) != 3 {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	id1 := strings.Split(r.URL.Path, "/")
	if len(id1) < 3 {
		renderError(w, http.StatusBadRequest, "ID not found")
		return
	}
	id := id1[len(id1)-1]

	relationTemplate, err := template.ParseFiles("template/relation.html")
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Error loading template: "+err.Error())
		return
	}
	baseURL := "https://groupietrackers.herokuapp.com/api/relation/"

	relations, err := ReadRelations(baseURL, id)
	if err != nil {

		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	if relations.ID == 0 {
		renderError(w, http.StatusNotFound, "Oops! We Can't find that page")
		return
	}

	err = relationTemplate.Execute(w, relations)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Failed to render template: "+err.Error())
	}
}
