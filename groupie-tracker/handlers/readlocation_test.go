package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadLocation(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/1" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":1,"locations":["north_carolina-usa","georgia-usa","los_angeles-usa","saitama-japan","osaka-japan","nagoya-japan","penrose-new_zealand","dunedin-new_zealand"]}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tests := []struct {
		name          string
		id            string
		wantID        int
		wantLocations []string
		wantErr       bool
	}{
		{
			name:          "Valid location",
			id:            "1",
			wantID:        1,
			wantLocations: []string{"north_carolina-usa", "georgia-usa", "los_angeles-usa", "saitama-japan", "osaka-japan", "nagoya-japan", "penrose-new_zealand", "dunedin-new_zealand"},
			wantErr:       false,
		},
		{
			name:          "Invalid location",
			id:            "999",
			wantID:        0,
			wantLocations: nil,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLocation(server.URL+"/", tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if int(got.ID) != tt.wantID {
				t.Errorf("ReadLocation() got ID = %v, want %v", got.ID, tt.wantID)
			}

			if !compareSlices(got.Locations, tt.wantLocations) {
				t.Errorf("ReadLocation() got Locations = %v, want %v", got.Locations, tt.wantLocations)
			}
		})
	}
}

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
