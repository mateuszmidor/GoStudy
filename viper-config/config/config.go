package config

import (
	"bytes"
	"os"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	InputData    InputData     `mapstructure:"InputData"`
	ReadInterval time.Duration `mapstructure:"ReadInterval"`
}

// InputData describes the application input data source
type InputData struct {
	GoogleDrive InputGoogleDrive `mapstructure:"GoogleDrive"`
}

// InputGoogleDrive describes Google Drive as the application input
type InputGoogleDrive struct {
	Email           string `mapstructure:"Email"`
	PrivateKey      string `mapstructure:"PrivateKey"`
	FilenamePattern string `mapstructure:"FilenamePattern"`
}

// Format enumerates supported configuration input formats
type Format string

const (
	// FormatYAML represents application configuration provided as YAML
	FormatYAML Format = "yaml"

	// FormatJSON represents application configuration provided as JSON
	FormatJSON Format = "json"
)

// LoadFromFile is a helper function that enables loading the configuration from file
func LoadFromFile(path string, format Format) (cfg Config, err error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	return LoadFromBytes(input, format)
}

// LoadFromBytes enables loading the configuration from plain text
func LoadFromBytes(input []byte, format Format) (cfg Config, err error) {
	viper.SetConfigType(string(format))
	if err := viper.ReadConfig(bytes.NewBuffer(input)); err != nil {
		return Config{}, err
	}

	// overwrite with env variables, if present
	loadEnvVariables()

	// overwrite with command line parameters, if present
	loadCmdLineParameters()

	// populate config struct
	err = viper.Unmarshal(&cfg)
	return
}

func loadEnvVariables() {
	viper.BindEnv("InputData.GoogleDrive.Email", "INPUT_GD_EMAIL")
	viper.BindEnv("InputData.GoogleDrive.PrivateKey", "INPUT_GD_PRIVATE_KEY")
	viper.BindEnv("InputData.GoogleDrive.FilenamePattern", "INPUT_GD_FILENAME_PATTERN")

	viper.BindEnv("ReadInterval", "READ_INTERVAL")

	viper.AutomaticEnv()
}

func loadCmdLineParameters() {
	pflag.String("InputData.GoogleDrive.Email", "", "input service account email")
	pflag.String("InputData.GoogleDrive.PrivateKey", "", "input service account private key")
	pflag.String("InputData.GoogleDrive.FilenamePattern", "", "pattern in the filename to look for when collecting files to read (simple find in string; not regex)")

	pflag.String("ReadInterval", "", "how often to read the data, e.g. 24h0m0s")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}
