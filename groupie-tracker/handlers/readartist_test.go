package api

import (
	"testing"
)

func TestReadArtist(t *testing.T) {
	tests := []struct {
		name            string
		id              string
		wantName        string
		wantMemberCount int
		wantErr         bool
	}{
		{
			name:            "Successful case",
			id:              "1",
			wantName:        "Queen",
			wantMemberCount: 7,
			wantErr:         false,
		},
		{
			name:            "API error",
			id:              "999",
			wantName:        "",
			wantMemberCount: 0,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL := "https://groupietrackers.herokuapp.com/api/artists/"
			got, err := ReadArtist(baseURL, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadArtist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.Name != tt.wantName {
					t.Errorf("ReadArtist() got Name = %v, want %v", got.Name, tt.wantName)
				}
				if len(got.Members) != tt.wantMemberCount {
					t.Errorf("ReadArtist() got %v members, want %v", len(got.Members), tt.wantMemberCount)
				}
			}
		})
	}
}
