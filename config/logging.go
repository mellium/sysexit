package config

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
