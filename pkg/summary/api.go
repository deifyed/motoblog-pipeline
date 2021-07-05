package summary

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/spf13/afero"
)

func GenerateTrips(fs *afero.Afero, routePaths []string, imagePaths []string, _ []string) ([]Trip, error) {
	trips := make([]Trip, len(routePaths))

	log.Debug().Strs("paths", routePaths).Msg("handling tracks")

	for index, path := range routePaths {
		result, err := handleRoute(fs, path)
		if err != nil {
			return nil, fmt.Errorf("handling track path: %w", err)
		}

		trips[index] = result
	}

	log.Debug().Strs("paths", imagePaths).Msg("handling images")

	for _, path := range imagePaths {
		err := handleImage(trips, fs, path)
		if err != nil {
			return nil, fmt.Errorf("handling image path: %w", err)
		}
	}

	return trips, nil
}
