package filesystem

import "github.com/spf13/afero"

type filesystemExporter struct {
	fs        *afero.Afero
	outputDir string
}

