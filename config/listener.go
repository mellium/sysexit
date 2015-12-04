package config

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
