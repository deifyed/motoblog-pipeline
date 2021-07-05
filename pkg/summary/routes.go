package summary

import (
	"fmt"
	"io"

	"github.com/hashicorp/go-uuid"

	"github.com/spf13/afero"
	"github.com/tkrajina/gpxgo/gpx"
)

func handleRoute(fs *afero.Afero, path string) (Trip, error) {
	routeFile, err := fs.Open(path)
	if err != nil {
		return Trip{}, fmt.Errorf("opening route file: %w", err)
	}

	defer func() {
		_ = routeFile.Close()
	}()

	content, err := io.ReadAll(routeFile)
	if err != nil {
		return Trip{}, fmt.Errorf("reading route file: %w", err)
	}

	gpxFile, err := gpx.ParseBytes(content)
	if err != nil {
		return Trip{}, fmt.Errorf("parsing gpx contents: %w", err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return Trip{}, fmt.Errorf("generating ID for trip: %w", err)
	}

	timeBounds := gpxFile.Tracks[0].TimeBounds()

	trip := Trip{
		ID:   id,
		From: timeBounds.StartTime,
		To:   timeBounds.EndTime,
		Tracks: []TripTrack{
			{
				SourcePath: path,
			},
		},
	}

	return trip, nil
}
