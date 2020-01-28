package main

type ArtistDTO struct {
	Topartists Topartists `json:"topartists"`
	Pointer    *string    `json:"-"` // ignore the rest of the json
}
type Artist struct {
	Name    string  `json:"name"`
	Pointer *string `json:"-"` // ignore the rest of the json
}
type Topartists struct {
	Artist  []Artist `json:"artist"`
	Pointer *string  `json:"-"` // ignore the rest of the json
}
