package concertutils

//JSONResponseA aa
type JSONResponse struct {
	ResultsPage ResultsPage `json:"resultsPage"`
	Pointer     *string     `json:"-"` // ignore the rest of the json
}
type Start struct {
	Date    string  `json:"date"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
type Performance struct {
	Artist  string  `json:"displayName"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
type Location struct {
	City    string  `json:"city"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
type Venue struct {
	Name    string  `json:"displayName"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
type Event struct {
	Name        string        `json:"displayName"`
	Start       Start         `json:"start"`
	Performance []Performance `json:"performance"`
	Venue       Venue         `json:"venue"`
	Location    Location      `json:"location"`
	Pointer     *string       `json:"-"` // ignore the rest of the json
}
type Results struct {
	Event   []Event `json:"event"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
type ResultsPage struct {
	Status  string  `json:"status"`
	Results Results `json:"results"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
