package honey

import (
	"github.com/BurntSushi/toml"
)

var defaultConfig = `
[[log]]

	level   = "info"
	console = true
`

type syslog struct {
	Network string `toml:"network"`
	RAddr   string `toml:"raddr"`
}

type log struct {
	Level    string `toml:"level"`
	Filename string `toml:"filename"`
	Console  bool   `toml:"console"`
	Syslog   syslog `toml:"syslog"`
}

type config struct {
	Loggers  []log `toml:"log"`
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

// DefaultConfig loads the default config for the application.
func DefaultConfig() *config {
	c := new(config)
	err := c.LoadBlob(defaultConfig)
	if err != nil {
		panic(err)
	}
	return c
}

// FromBlob loads a new config struct from the given TOML blob.
func FromBlob(blob string) (c *config, err error) {
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
