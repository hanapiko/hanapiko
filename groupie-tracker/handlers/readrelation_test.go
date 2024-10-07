package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchRelations(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/1":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":1,"datesLocations":{"new_york":["2019-10-05","2019-10-06"],"london":["2019-11-20","2019-11-21"]}}`))
		case "/2":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":2,"datesLocations":{"paris":["2020-01-01"],"berlin":["2020-02-02","2020-02-03"]}}`))
		case "/error":
			w.WriteHeader(http.StatusInternalServerError)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tests := []struct {
		name    string
		id      string
		wantID  int64
		wantErr bool
	}{
		{"Valid relation 1", "1", 1, false},
		{"Valid relation 2", "2", 2, false},
		{"Invalid relation", "999", 0, true},
		{"Server error", "error", 0, true},
		{"Empty ID", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadRelations(server.URL+"/", tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchRelations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID != tt.wantID {
				t.Errorf("FetchRelations() got ID = %v, want %v", got.ID, tt.wantID)
			}

			if !tt.wantErr && (len(got.Locations) == 0) {
				t.Errorf("FetchRelations() got empty DatesLocations, expected some data")
			}
		})
	}
}
