package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/mateuszmidor/GoStudy/viper-config/config"
	"github.com/stretchr/testify/assert"
)

var INPUT_CONFIG = []byte(`
InputData:
  GoogleDrive:
    Email: example_input_gd_email
    PrivateKey: example_input_gd_private_key
    FilenamePattern: example_pattern 

ReadInterval: 24h0m0s  
`)

func Test_EnvVars_Override_ConfigFile(t *testing.T) {
	// given
	const INPUT_GD_EMAIL = "input-email@acme.com"
	const INPUT_GD_PRIVATE_KEY = "input-private-key"
	const INPUT_GD_FILENAME_PATTERN = "input-file-pattern"
	const READ_INTERVAL = time.Hour * 24

	// when
	os.Setenv("INPUT_GD_EMAIL", INPUT_GD_EMAIL)
	os.Setenv("INPUT_GD_PRIVATE_KEY", INPUT_GD_PRIVATE_KEY)
	os.Setenv("INPUT_GD_FILENAME_PATTERN", INPUT_GD_FILENAME_PATTERN)
	os.Setenv("READ_INTERVAL", READ_INTERVAL.String())
	cfg, err := config.LoadFromBytes(INPUT_CONFIG, config.FormatYAML)

	// then
	assert := assert.New(t)
	assert.NoError(err)

	assert.Equal(INPUT_GD_EMAIL, cfg.InputData.GoogleDrive.Email)
	assert.Equal(INPUT_GD_PRIVATE_KEY, cfg.InputData.GoogleDrive.PrivateKey)
	assert.Equal(INPUT_GD_FILENAME_PATTERN, cfg.InputData.GoogleDrive.FilenamePattern)
	assert.Equal(READ_INTERVAL, cfg.ReadInterval)
}

func Test_CmdLineParams_Override_EnvVars(t *testing.T) {
	// TODO: how to inject cmd line params when in unit test?
}
