package concertutils

import "time"

type Concert struct {
	Event  string
	Date   time.Time
	Artist string
	Venue  string
	City   string
}

func CreateConcert(event string, date time.Time, artist string, venue string, city string) *Concert {
	return &Concert{
		Event:  event,
		Date:   date,
		Artist: artist,
		Venue:  venue,
		City:   city,
	}
}
