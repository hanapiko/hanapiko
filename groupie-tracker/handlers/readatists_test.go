package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadArtists(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"name":"Artist1"},{"id":2,"name":"Artist2"}]`))
	}))
	defer server.Close()

	// Use the mock server URL for testing
	testURL := server.URL

	// Test the ReadArtists function with the test URL
	artists, err := ReadArtists(testURL)
	if err != nil {
		t.Fatalf("ReadArtists() returned an error: %v", err)
	}

	// Check the number of artists
	if len(artists) != 2 {
		t.Errorf("Expected 2 artists, got %d", len(artists))
	}

	// Check the content of the first artist
	if artists[0].ID != 1 || artists[0].Name != "Artist1" {
		t.Errorf("First artist does not match expected values")
	}

	// Check the content of the second artist
	if artists[1].ID != 2 || artists[1].Name != "Artist2" {
		t.Errorf("Second artist does not match expected values")
	}
}
