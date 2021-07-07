package main

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/deifyed/motoblog/pkg/exporter/filesystem"

	"github.com/deifyed/motoblog/pkg/config"
	"github.com/deifyed/motoblog/pkg/summary"
	"github.com/deifyed/motoblog/pkg/walker"
	"github.com/spf13/afero"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Errorf("loading config: %w", err))
	}

	err = cfg.Validate()
	if err != nil {
		panic(fmt.Errorf("validating config: %w", err))
	}

	fs := &afero.Afero{Fs: afero.NewOsFs()}

	gatheredPaths, err := gatherPaths(fs, cfg)
	if err != nil {
		log.Error().Err(err).Msg("gathering relevant paths")

		return
	}

	trips, err := summary.GenerateTrips(fs, gatheredPaths.tracks, gatheredPaths.images, gatheredPaths.notes)
	if err != nil {
		log.Error().Err(err).Msg("generating trips")

		return
	}

	tripExporter := filesystem.NewFilesystemExporter(fs, cfg.DestinationDir)

	for _, trip := range trips {
		err = tripExporter.Export(trip)
		if err != nil {
			log.Error().Err(err).Msg("exporting trip")

			return
		}
	}
}

type paths struct {
	tracks []string
	images []string
	notes  []string
}

func gatherPaths(fs *afero.Afero, cfg config.Config) (paths, error) {
	images, err := walker.Walk(walker.WalkOpts{
		Fs:         fs,
		SourceDir:  cfg.ImagesSourceDir,
		Extensions: []string{".jpeg", ".jpg", ".png"},
	})
	if err != nil {
		return paths{}, fmt.Errorf("fetching images: %w", err)
	}

	tracks, err := walker.Walk(walker.WalkOpts{
		Fs:         fs,
		SourceDir:  cfg.TracksSourceDir,
		Extensions: []string{".gpx"},
	})
	if err != nil {
		return paths{}, fmt.Errorf("fetching routes: %w", err)
	}

	notes, err := walker.Walk(walker.WalkOpts{
		Fs:         fs,
		SourceDir:  cfg.NotesSourceDir,
		Extensions: []string{".md", ".txt"},
	})
	if err != nil {
		return paths{}, fmt.Errorf("fetching notes: %w", err)
	}

	return paths{
		tracks: tracks,
		images: images,
		notes:  notes,
	}, nil
}
