package filesystem

import (
	"fmt"
	"os"
	"path"

	"github.com/deifyed/motoblog/pkg/exporter"
	"github.com/deifyed/motoblog/pkg/summary"
	"github.com/spf13/afero"
)

func (f filesystemExporter) Export(trip summary.Trip) error {
	tripOutputDir := path.Join(f.outputDir, trip.ID)

	err := f.fs.MkdirAll(tripOutputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("creating trip directory: %w", err)
	}

	for index, routeTrackFile := range trip.Tracks {
		destinationFilePath := path.Join(
			tripOutputDir,
			fmt.Sprintf("track-%d%s", index+1, path.Ext(routeTrackFile.SourcePath)),
		)

		err := copyFile(f.fs, routeTrackFile.SourcePath, destinationFilePath)
		if err != nil {
			return fmt.Errorf("copying track file: %w", err)
		}
	}

	for index, imageFile := range trip.Images {
		destinationFilePath := path.Join(
			tripOutputDir,
			fmt.Sprintf("image-%d%s", index+1, path.Ext(imageFile.SourcePath)),
		)

		err := copyFile(f.fs, imageFile.SourcePath, destinationFilePath)
		if err != nil {
			return fmt.Errorf("copying image file: %w", err)
		}
	}

	return nil
}

func NewFilesystemExporter(fs *afero.Afero, outputDir string) exporter.Interface {
	return &filesystemExporter{
		fs:        fs,
		outputDir: outputDir,
	}
}
