package honey

import (
	"github.com/BurntSushi/toml"
)

var defaultConfig = []byte(`
[[c2s.listen]]
addr = "test"

[[log]]

	level    = "info"

	[log.syslog]
	network = ""
	raddr = ""
`)

// A Service is anything for which we should spawn a listener.
type Service struct {
	Addr     string `toml:"addr"`
	Certfile string `toml:"certfile"`
	Keyfile  string `toml:"keyfile"`
}

// Host is a host for which we'll route XMPP stanzas.
type Host struct {
	Name     string `toml:"name"`
	Certfile string `toml:"certfile"`
	Keyfile  string `toml:"keyfile"`
}

// Hosts is a list of all the hosts served.
type Hosts []Host

// C2S containst configuration for client to server connections.
type C2S struct {
	Services []Service `toml:"listen"`
}

// Syslog describes a local (empty network string) or remote syslog server.
type Syslog struct {
	Network string `toml:"network"`
	RAddr   string `toml:"raddr"`
}

// Log represents a single logger that may log to a syslog server, a filename,
// and/or the console.
type Log struct {
	Level    string `toml:"level"`
	Filename string `toml:"filename"`
	Console  bool   `toml:"console"`
	Syslog   Syslog `toml:"syslog"`
}

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
