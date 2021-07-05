package filesystem

import (
	"fmt"

	"github.com/spf13/afero"
)

func copyFile(fs *afero.Afero, fromPath, toPath string) error {
	content, err := fs.ReadFile(fromPath)
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	err = fs.WriteFile(toPath, content, 0o744)
	if err != nil {
		return fmt.Errorf("writing track file to destination dir: %w", err)
	}

	return nil
}
