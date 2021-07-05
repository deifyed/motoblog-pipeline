package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/spf13/viper"
)

func Load() (Config, error) {
	viper.AutomaticEnv()

	return Config{
		DestinationDir:  viper.GetString("DESTINATION_DIR"),
		ImagesSourceDir: viper.GetString("IMAGES_SOURCE_DIR"),
		TracksSourceDir: viper.GetString("TRACKS_SOURCE_DIR"),
		NotesSourceDir: viper.GetString("NOTES_SOURCE_DIR"),
	}, nil
}

func (receiver Config) Validate() error {
	return validation.ValidateStruct(&receiver,
		validation.Field(&receiver.DestinationDir, validation.Required),
		validation.Field(&receiver.TracksSourceDir, validation.Required),
	)
}
