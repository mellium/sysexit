package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/SamWhited/koine"
	"github.com/SamWhited/logger"
)

var C *config

func init() {
	C = new(config)
}

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

// S2S configuration.
type s2s struct {
	Services []service `toml:"listen"`
}

// C2S configuration.
type c2s struct {
	Services []service `toml:"listen"`
}

// Configuration struct.
type config struct {
	S2S   s2s   `toml:"s2s"`
	C2S   c2s   `toml:c2s"`
	Hosts hosts `toml:"host"`
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
func ReloadConfig() {
	if info, err := os.Stat(CONFIG_FILE); err == nil && !info.IsDir() {
		if err := C.LoadFile(CONFIG_FILE); err != nil {
			logger.Err("Error parsing config file " + CONFIG_FILE + ": " + err.Error())
			return
		}
	}
}
