package config

import (
	"github.com/BurntSushi/toml"
)

// config holds the entire application configuration.
type config struct {
	C2S     C2S
	Loggers []Log `toml:"log"`
}

// LoadBlob parses a TOML blob and unmarshals it into the config struct.
func (c *config) LoadBlob(blob []byte) error {
	return toml.Unmarshal(blob, c)
}

// LoadFile loads a TOML file and unmarshals it into the config struct.
func (c *config) LoadFile(path string) error {
	_, err := toml.DecodeFile(path, c)
	return err
}

// Default loads the default config for the application.
func Default() *config {
	c := new(config)
	err := c.LoadBlob(defaultConfig)
	if err != nil {
		panic(err)
	}
	return c
}

// FromBlob loads a new config struct from the given TOML blob.
func FromBlob(blob []byte) (c *config, err error) {
	c = new(config)
	err = c.LoadBlob(blob)
	return
}

// FromFile loads a new config struct from the given TOML file.
func FromFile(path string) (c *config, err error) {
	c = new(config)
	err = c.LoadFile(path)
	return
}
