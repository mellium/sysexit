package config

import (
	"io"

	"github.com/BurntSushi/toml"
)

// Config represents a fully parsed configuration.
type Config struct {
	Log      Log           `toml:"log"`
	VHosts   []VirtualHost `toml:"virtualhost"`
	metaData toml.MetaData `toml:"-"`
}

// Log configures the output logging.
// Errors are always logged and cannot be turned off.
// Verbose tells the server to log more information on errors if possible and to
// log non-error messages to stdout.
type Log struct {
	Verbose bool
}

// C2S contains configuration for client to server connections.
type VirtualHost struct {
	Domain string  `toml:"domain"`
	C2S    Service `toml:"c2s"`
	S2S    Service `toml:"s2s"`
}

// A Service is anything for which we should spawn a listener.
type Service struct {
	Addr     string `toml:"addr"`
	Certfile string `toml:"certfile"`
	Keyfile  string `toml:"keyfile"`
	DHParams string `toml:"dhparams"`
}

// LoadBlob parses a TOML blob and unmarshals it into the config struct.
func (c *Config) LoadBlob(blob string) error {
	md, err := toml.Decode(blob, c)
	c.metaData = md
	return err
}

// LoadFile loads a TOML file and unmarshals it into the config struct.
func (c *Config) LoadFile(path string) error {
	md, err := toml.DecodeFile(path, c)
	c.metaData = md
	return err
}

// LoadReader decodes TOML from r into the config struct.
func (c *Config) LoadReader(r io.Reader) error {
	md, err := toml.DecodeReader(r, c)
	c.metaData = md
	return err
}

// MetaData allows you to access information about the TOML keys that were
// parsed to generate the config.
func (c *Config) MetaData() toml.MetaData {
	return c.metaData
}

// FromBlob loads a new config struct from the given TOML blob.
func FromBlob(blob string) (c *Config, err error) {
	c = new(Config)
	err = c.LoadBlob(blob)
	return
}

// FromFile loads a new config struct from the given TOML file.
func FromFile(path string) (c *Config, err error) {
	c = new(Config)
	err = c.LoadFile(path)
	return
}

// FromReader parses TOML from r into c.
func FromReader(r io.Reader) (c *Config, err error) {
	c = new(Config)
	err = c.LoadReader(r)
	return
}
