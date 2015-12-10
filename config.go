package honey

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Log      Log           `toml:"log"`
	VHosts   []VirtualHost `toml:"virtualhost"`
	metaData toml.MetaData
}

// LoadBlob parses a TOML blob and unmarshals it into the config struct.
func (c *config) LoadBlob(blob string) error {
	md, err := toml.Decode(blob, c)
	c.metaData = md
	return err
}

// LoadFile loads a TOML file and unmarshals it into the config struct.
func (c *config) LoadFile(path string) error {
	md, err := toml.DecodeFile(path, c)
	c.metaData = md
	return err
}

// MetaData allows you to access information about the TOML keys that were
// parsed to generate the config.
func (c *config) MetaData() toml.MetaData {
	return c.metaData
}

// ConfigFromBlob loads a new config struct from the given TOML blob.
func ConfigFromBlob(blob string) (c *config, err error) {
	c = new(config)
	err = c.LoadBlob(blob)
	return
}

// ConfigFromFile loads a new config struct from the given TOML file.
func ConfigFromFile(path string) (c *config, err error) {
	c = new(config)
	err = c.LoadFile(path)
	return
}
