package summary

import "time"

type Trip struct {
	ID string

	From time.Time
	To   time.Time

	Tracks []TripTrack
	Texts  []TripText
	Images []TripImage
}

type TripTrack struct {
	SourcePath string
}

type TripText struct {
	SourcePath string
}

type TripImage struct {
	SourcePath string
}
