package walker

import (
	"fmt"
	"io/fs"
	path2 "path"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/spf13/afero"
)

type WalkOpts struct {
	Fs         *afero.Afero
	SourceDir  string
	Extensions []string
}

func Walk(opts WalkOpts) ([]string, error) {
	paths := make([]string, 0)
	validExtensions := map[string]int{}

	for _, ext := range opts.Extensions {
		validExtensions[strings.ToUpper(ext)] = 0
		validExtensions[strings.ToLower(ext)] = 0
	}

	e := log.Debug().Strs("extensions", opts.Extensions)

	err := opts.Fs.Walk(opts.SourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("walking dir: %w", err)
		}

		currentExtension := path2.Ext(path)
		_, ok := validExtensions[currentExtension]

		e.Str("current", currentExtension)
		e.Bool("matching", ok)

		if !ok {
			e.Send()

			return nil
		}

		validExtensions[currentExtension]++

		e.Msg("relevant")

		paths = append(paths, path)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking directory: %w", err)
	}

	return paths, nil
}
