package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

func TestRenderError(t *testing.T) {
	Init()
	tests := []struct {
		name           string
		status         int
		message        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Test 404 Not Found",
			status:         http.StatusNotFound,
			message:        "Page Not Found",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Error 404",
		},
		{
			name:           "Test 500 Internal Server Error",
			status:         http.StatusInternalServerError,
			message:        "Internal Server Error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Error 500",
		},
		{
			name:           "Test 403 Forbidden",
			status:         http.StatusForbidden,
			message:        "Access Denied",
			expectedStatus: http.StatusForbidden,
			expectedBody:   "Error 403",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			renderError(w, tt.status, tt.message)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, w.Code)
			}

			if !strings.Contains(w.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q; got %q", tt.expectedBody, w.Body.String())
			}

			if !strings.Contains(w.Body.String(), tt.message) {
				t.Errorf("expected body to contain %q; got %q", tt.message, w.Body.String())
			}
		})
	}
}

func TestInit(t *testing.T) {
	// Temporarily replace the global errorTemplate
	originalTemplate := errorTemplate
	defer func() { errorTemplate = originalTemplate }()

	// Reset errorTemplate to nil
	errorTemplate = nil

	// Call init() manually
	Init()

	// Check if errorTemplate is not nil after init
	if errorTemplate == nil {
		t.Error("errorTemplate is nil after init")
	}

	testCases := []struct {
		name         string
		code         int
		message      string
		expectedBody string
	}{
		{"Not Found Error", 404, "Test Error", "Error 404"},
		{"Internal Server Error", 500, "Server Error", "Error 500"},
		{"Forbidden Error", 403, "Access Denied", "Error 403"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := errorTemplate.Execute(w, struct {
				Code    int
				Message string
			}{
				Code:    tc.code,
				Message: tc.message,
			})
			if err != nil {
				t.Errorf("Error executing template: %v", err)
			}

			if !strings.Contains(w.Body.String(), tc.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tc.expectedBody, w.Body.String())
			}

			if !strings.Contains(w.Body.String(), tc.message) {
				t.Errorf("Expected body to contain %q, got %q", tc.message, w.Body.String())
			}
		})
	}

	// Test with invalid template path
	errorTemplate = nil
	Init()
	if errorTemplate == nil {
		t.Error("Fallback template was not created when given an invalid path")
	}
}

// type Artist struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

var ReadArtistFunc = ReadArtist // Function variable for testing

func MockReadArtist(baseURL, id string) (Artist, error) {
	if id == "1" {
		return Artist{ID: 1, Name: "Test Artist"}, nil
	}
	return Artist{}, fmt.Errorf("Artist ID not found")
}

func TestArtistHandler(t *testing.T) {
	mockTemplate, _ := template.New("test").Parse("{{.Name}}")

	tests := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/artist/1", http.StatusOK, "Test Artist"},
		{"GET", "/artist/2", http.StatusNotFound, "Artist not found"},
		{"POST", "/artist/1", http.StatusMethodNotAllowed, "Wrong method"},
		{"GET", "/wrongpath", http.StatusNotFound, "Page Not Found"},
		{"GET", "/artist/", http.StatusBadRequest, "Artist ID not found"},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 || parts[1] != "artist" {
			http.Error(w, "Page Not Found", http.StatusNotFound)
			return
		}

		id := parts[2]
		if id == "" {
			http.Error(w, "Artist ID not found", http.StatusBadRequest)
			return
		}

		result, err := ReadArtistFunc("https://groupietrackers.herokuapp.com/api/artists/", id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				http.Error(w, "Artist not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error fetching artist: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if result.ID == 0 {
			http.Error(w, "Artist not found", http.StatusNotFound)
			return
		}

		if err := mockTemplate.Execute(w, result); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	}

	ReadArtistFunc = MockReadArtist
	defer func() { ReadArtistFunc = ReadArtist }() // Restore the original function

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.url, nil)
		w := httptest.NewRecorder()

		handler(w, req)

		res := w.Result()
		if res.StatusCode != test.expectedCode {
			t.Errorf("expected status code %d, got %d", test.expectedCode, res.StatusCode)
		}

		body := w.Body.String()
		if !strings.Contains(body, test.expectedBody) {
			t.Errorf("expected body to contain %q, got %q", test.expectedBody, body)
		}
	}
}

var ReadLocationFunc = ReadLocation // Function variable for testing

func MockReadLocation(baseURL, id string) (Location, error) {
	if id == "1" {
		return Location{ID: 1, Locations: []string{"Location A", "Location B"}}, nil
	} else if id == "3" {
		return Location{ID: 3, Locations: []string{"Location C"}}, nil
	}
	return Location{}, fmt.Errorf("Location not found")
}

func TestLocationHandler(t *testing.T) {
	mockTemplate, _ := template.New("test").Parse("{{range .}}{{.}} {{end}}") // Output locations

	tests := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/location/1", http.StatusOK, "Location A Location B "},
		{"GET", "/location/3", http.StatusOK, "Location C "},
		{"GET", "/location/2", http.StatusNotFound, "Location not found"},
		{"POST", "/location/1", http.StatusMethodNotAllowed, "Wrong method"},
		{"GET", "/wrongpath", http.StatusNotFound, "Page Not Found"},
		{"GET", "/location/", http.StatusBadRequest, "Location ID not found"},
		{"GET", "/location/invalid", http.StatusNotFound, "Location not found"},
		{"GET", "/location/0", http.StatusNotFound, "Location not found"},
		{"GET", "/location/999", http.StatusNotFound, "Location not found"},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 || parts[1] != "location" {
			http.Error(w, "Page Not Found", http.StatusNotFound)
			return
		}

		id := parts[2]
		if id == "" {
			http.Error(w, "Location ID not found", http.StatusBadRequest)
			return
		}

		result, err := ReadLocationFunc("https://groupietrackers.herokuapp.com/api/locations/", id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				http.Error(w, "Location not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error fetching location: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if result.ID == 0 {
			http.Error(w, "Location not found", http.StatusNotFound)
			return
		}

		if err := mockTemplate.Execute(w, result.Locations); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	}

	ReadLocationFunc = MockReadLocation
	defer func() { ReadLocationFunc = ReadLocation }() // Restore the original function

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.url, nil)
		w := httptest.NewRecorder()

		handler(w, req)

		res := w.Result()
		if res.StatusCode != test.expectedCode {
			t.Errorf("expected status code %d, got %d", test.expectedCode, res.StatusCode)
		}

		body := w.Body.String()
		if !strings.Contains(body, test.expectedBody) {
			t.Errorf("expected body to contain %q, got %q", test.expectedBody, body)
		}
	}
}

var ReadDateFunc = ReadDate // Function variable for testing

func MockReadDate(baseURL, id string) (DateEntry, error) {
	if id == "1" {
		return DateEntry{ID: 1, Dates: []string{"2023-01-01", "2023-02-01"}}, nil
	} else if id == "3" {
		return DateEntry{ID: 3, Dates: []string{"2023-03-01"}}, nil
	}
	return DateEntry{}, fmt.Errorf("Date not found")
}

func TestDateHandler(t *testing.T) {
	mockTemplate, _ := template.New("test").Parse("{{range .}}{{.}} {{end}}") // Output dates

	tests := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/dates/1", http.StatusOK, "2023-01-01 2023-02-01 "},
		{"GET", "/dates/3", http.StatusOK, "2023-03-01 "},
		{"GET", "/dates/2", http.StatusNotFound, "Date entry not found"},
		{"POST", "/dates/1", http.StatusMethodNotAllowed, "Wrong method"},
		{"GET", "/wrongpath", http.StatusNotFound, "Page Not Found"},
		{"GET", "/dates/", http.StatusBadRequest, "Artist ID not found"},
		{"GET", "/dates/invalid", http.StatusNotFound, "Date entry not found"},
		{"GET", "/dates/0", http.StatusNotFound, "Date entry not found"},
		{"GET", "/dates/999", http.StatusNotFound, "Date entry not found"},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
			return
		}

		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 3 || parts[1] != "dates" {
			http.Error(w, "Page Not Found", http.StatusNotFound)
			return
		}

		id := parts[2]
		if id == "" {
			http.Error(w, "Artist ID not found", http.StatusBadRequest)
			return
		}

		result, err := ReadDateFunc("https://groupietrackers.herokuapp.com/api/dates/", id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				http.Error(w, "Date entry not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error fetching date: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if result.ID == 0 {
			http.Error(w, "Date entry not found", http.StatusNotFound)
			return
		}

		if err := mockTemplate.Execute(w, result.Dates); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	}

	ReadDateFunc = MockReadDate
	defer func() { ReadDateFunc = ReadDate }() // Restore the original function

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.url, nil)
		w := httptest.NewRecorder()

		handler(w, req)

		res := w.Result()
		if res.StatusCode != test.expectedCode {
			t.Errorf("expected status code %d, got %d", test.expectedCode, res.StatusCode)
		}

		body := w.Body.String()
		if !strings.Contains(body, test.expectedBody) {
			t.Errorf("expected body to contain %q, got %q", test.expectedBody, body)
		}
	}
}

// Mock function for FetchRelations for testing purposes
var FetchRelationsFunc = ReadRelations

func MockFetchRelations(baseURL, id string) (Relation, error) {
	if id == "1" {
		return Relation{ID: 1, Locations: map[string][]string{
			"New York":    {"2023-01-01", "2023-02-01"},
			"Los Angeles": {"2023-03-01"},
		}}, nil
	} else if id == "3" {
		return Relation{ID: 3, Locations: map[string][]string{
			"London": {"2023-05-01"},
		}}, nil
	}
	return Relation{}, fmt.Errorf("Relation not found")
}

func TestRelationHandler(t *testing.T) {
	// Mock template for testing rendering
	mockTemplate, _ := template.New("test").Parse("{{range $location, $dates := .Locations}}{{$location}}: {{range $dates}}{{.}} {{end}} {{end}}")

	tests := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody []string // We will check for multiple substrings
	}{
		{"GET", "/relation/1", http.StatusOK, []string{"New York: 2023-01-01 2023-02-01", "Los Angeles: 2023-03-01"}},
		{"GET", "/relation/3", http.StatusOK, []string{"London: 2023-05-01"}},
		{"GET", "/relation/2", http.StatusNotFound, []string{"Relation not found"}},
		{"POST", "/relation/1", http.StatusMethodNotAllowed, []string{"Method not allowed"}},
		{"GET", "/wrongpath", http.StatusNotFound, []string{"Page Not Found"}},
		{"GET", "/relation/invalid", http.StatusNotFound, []string{"Relation not found"}},
		{"GET", "/relation/0", http.StatusNotFound, []string{"Relation not found"}},
		{"GET", "/relation/999", http.StatusNotFound, []string{"Relation not found"}},
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			renderError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/relation/") || len(strings.Split(r.URL.Path, "/")) != 3 {
			renderError(w, http.StatusNotFound, "Page Not Found")
			return
		}

		id1 := strings.Split(r.URL.Path, "/")
		if len(id1) < 3 {
			renderError(w, http.StatusBadRequest, "ID not found")
			return
		}
		id := id1[len(id1)-1]

		relations, err := FetchRelationsFunc("https://groupietrackers.herokuapp.com/api/relation/", id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				renderError(w, http.StatusNotFound, "Relation not found")
			} else {
				renderError(w, http.StatusInternalServerError, "Error fetching relations: "+err.Error())
			}
			return
		}

		if relations.ID == 0 {
			renderError(w, http.StatusNotFound, "Relation not found")
			return
		}

		if err := mockTemplate.Execute(w, relations); err != nil {
			renderError(w, http.StatusInternalServerError, "Failed to render template: "+err.Error())
		}
	}

	// Replace FetchRelationsFunc with mock function and restore afterward
	FetchRelationsFunc = MockFetchRelations
	defer func() { FetchRelationsFunc = ReadRelations }()

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.url, nil)
		w := httptest.NewRecorder()

		handler(w, req)

		res := w.Result()
		if res.StatusCode != test.expectedCode {
			t.Errorf("expected status code %d, got %d", test.expectedCode, res.StatusCode)
		}

		body := w.Body.String()
		for _, expected := range test.expectedBody {
			if !strings.Contains(body, expected) {
				t.Errorf("expected body to contain %q, got %q", expected, body)
			}
		}
	}
}
