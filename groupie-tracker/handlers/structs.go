package api

type Location struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
}
type DateEntry struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type Relation struct {
	ID        int64               `json:"id"`
	Locations map[string][]string `json:"datesLocations"`
}
