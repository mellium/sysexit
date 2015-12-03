package config

import (
	"../logging"

	"github.com/BurntSushi/toml"
)

// A service for which we should spawn a listener.
type service struct {
	Addr     string `toml:"addr"`
	Certfile string `toml:"certfile"`
	Keyfile  string `toml:"keyfile"`
}

// A host for which we'll route XMPP stanzas.
type host struct {
	Name     string `toml:"name"`
	Certfile string `toml:"certfile"`
	Keyfile  string `toml:"keyfile"`
}

// Hosts configuration.
type hosts []host

// C2S configuration.
type c2s struct {
	Services []service `toml:"listen"`
}

// Configuration struct.
type config struct {
	C2S     c2s            `toml:c2s"`
	Hosts   hosts          `toml:"host"`
	Logging logging.Config `toml:"logging"`
}

// Loads a TOML blob into the config struct.
func (c *config) LoadBlob(blob string) error {
	_, err := toml.Decode(blob, c)
	return err
}

// Loads a TOML file into the config struct.
func (c *config) LoadFile(path string) error {
	_, err := toml.DecodeFile(path, c)
	return err
}

// Config related utility functions

const CONFIG_FILE string = "honey.config"

// Flush config and reload.
func Load(path string) (*config, error) {
	c := new(config)
	_, err := toml.DecodeFile(path, c)
	return c, err
}
