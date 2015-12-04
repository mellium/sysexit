package config

var defaultConfig = []byte(`
[[c2s.listen]]
addr = "test"

[[log]]

	level    = "info"

	[log.syslog]
	network = ""
	raddr = ""
`)
