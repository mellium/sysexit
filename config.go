package mel

import (
	"sync/atomic"

	"github.com/BurntSushi/toml"
)

// Config represents a complete set of configuration values for the server. By
// default, a global instance of Config is created that is safe for concurrent
// access by many goroutines.
type Config struct {
	Log      Log           `toml:"log"`
	VHosts   []VirtualHost `toml:"virtualhost"`
	metaData toml.MetaData
}

// MetaData allows you to access information about the TOML keys that were
// parsed to generate the config.
func (c *Config) MetaData() toml.MetaData {
	return c.metaData
}

var config atomic.Value

// LoadBlob parses a TOML blob and unmarshals it into the global config struct.
func LoadBlob(blob string) (*Config, error) {
	var c *Config
	md, err := toml.Decode(blob, c)
	c.metaData = md
	config.Store(&c)
	return config.Load().(*Config), err
}

// LoadFile loads a TOML file and unmarshals it into the global config struct.
func LoadFile(path string) (*Config, error) {
	var c *Config
	md, err := toml.DecodeFile(path, c)
	c.metaData = md
	return config.Load().(*Config), err
}

func Load() *Config {
	return config.Load().(*Config)
}
