package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadDate(t *testing.T) {
	baseURL := "https://groupietrackers.herokuapp.com/api/dates/"
	tests := []struct {
		name           string
		id             string
		mockResponse   string
		mockStatusCode int
		wantErr        bool
	}{
		{
			name:           "Successful case",
			id:             "1",
			mockResponse:   `{"id": 1, "dates": ["2019-01-01", "2019-01-02"]}`,
			mockStatusCode: http.StatusOK,
			wantErr:        false,
		},
		{
			name:           "API error",
			id:             "999",
			mockResponse:   ``,
			mockStatusCode: http.StatusNotFound,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			// Replace the API URL with our mock server URL
			origURL := baseURL
			baseURL = server.URL + "/"
			got, err := ReadDate(baseURL, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.ID != 1 {
					t.Errorf("ReadDate() got ID = %v, want %v", got.ID, 1)
				}
				if len(got.Dates) != 2 {
					t.Errorf("ReadDate() got %v dates, want %v", len(got.Dates), 2)
				}
			}
			// fmt.Println(got.Dates)

			// Restore the original URL
			baseURL = origURL
		})
	}
}
