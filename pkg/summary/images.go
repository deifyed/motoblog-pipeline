package summary

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/evanoberholster/imagemeta/exif"
	"github.com/evanoberholster/imagemeta/meta"

	"github.com/rs/zerolog/log"

	"github.com/spf13/afero"

	"github.com/evanoberholster/imagemeta"
)

func handleImage(trips []Trip, fs *afero.Afero, path string) error {
	metadata, err := acquireMetadata(fs, path)
	if err != nil {
		if errors.Is(err, imagemeta.ErrNoExif) {
			log.Warn().Str("path", path).Err(err).Msg("missing exif data")

			return nil
		}

		return fmt.Errorf("acquiring metadata: %w", err)
	}

	for index, trip := range trips {
		e := log.Debug().Str("path", path)
		e.Str("trip", trip.ID)

		from := getStartOfDay(trip.From)
		to := getEndOfDay(trip.To)

		if !metadata.created.After(from) {
			e.Msg("image too old")

			continue
		}

		if !metadata.created.After(to) {
			e.Msg("image too recent")

			continue
		}

		e.Msg("matching trip")

		trips[index].Images = append(trips[index].Images, TripImage{SourcePath: path})
	}

	return nil
}

type imageMetadata struct {
	created time.Time
}

func acquireMetadata(fs *afero.Afero, path string) (imageMetadata, error) {
	imageFile, err := fs.Open(path)
	if err != nil {
		return imageMetadata{}, fmt.Errorf("opening file: %w", err)
	}

	var e *exif.Data

	exifDecodeFn := func(r io.Reader, m *meta.Metadata) error {
		e, _ = e.ParseExifWithMetadata(imageFile, m)
		return nil
	}
	xmpDecodeFn := func(_ io.Reader, _ *meta.Metadata) error {
		return nil
	}

	_, err = imagemeta.NewMetadata(imageFile, xmpDecodeFn, exifDecodeFn)
	if err != nil {
		return imageMetadata{}, fmt.Errorf("parsing metadata: %w", err)
	}

	created, err := e.DateTime()
	if err != nil {
		return imageMetadata{}, fmt.Errorf("extracting created time: %w", err)
	}

	return imageMetadata{
		created: created,
	}, nil
}

func getStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func getEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, t.Location())
}
