package logging

import (
	"bitbucket.org/SamWhited/multilogger"
	"log"
)

type syslogConfig struct {
	network string `toml:"network"`
	raddr   string `toml:"raddr"`
}

type severityLevelConfig struct {
	filename string `toml:"filename"`
	console  bool   `toml:"console"`
}

// The logging config format includes a table of severity levels which can
// include filename's and the option to also output to the console (severity
// levels of `Error` and below will log to STDERR, all others will log to
// STDOUT). It may also include information about a syslog server (the network
// and relay address) which should also receive log messages. An empty network
// and raddr indicates that we should use the local syslog server.
//
//    [logging]
//
//      [logging.err]
//      filename="/var/log/honey.err.log"
//      console=true
//
//      [logging.info]
//      filename="/var/log/honey.log"
//
//      [logging.syslog]
//      network="127.0.0.0"
//      raddr=""
//
type Config struct {
	emerg  severityLevelConfig `toml:"emerg"`
	alert  severityLevelConfig `toml:"alert"`
	crit   severityLevelConfig `toml:"crit"`
	err    severityLevelConfig `toml:"err"`
	warn   severityLevelConfig `toml:"warning"`
	notice severityLevelConfig `toml:"notice"`
	info   severityLevelConfig `toml:"info"`
	debug  severityLevelConfig `toml:"debug"`
	syslog syslogConfig        `toml:"syslog"`
}

// Log is represents a set of multiloggers which will be used for logging
// events to the provided loggers.
type Log struct {
	config  Config
	console *multilogger.MultiLogger
	emerg   *multilogger.MultiLogger
	alert   *multilogger.MultiLogger
	crit    *multilogger.MultiLogger
	err     *multilogger.MultiLogger
	warn    *multilogger.MultiLogger
	notice  *multilogger.MultiLogger
	info    *multilogger.MultiLogger
	debug   *multilogger.MultiLogger
}

// New creates a new logger that will obey the given configuration.
func New(config Config) (*Log, error) {
	l := &Log{
		config: config,
	}
	emerglog := []log.Logger{}
	if l.config.emerg.filename != "" {
	}
	if l.config.emerg.filename != "" || l.config.emerg.console {
		l.emerg = multilogger.New()
	}
	return l, nil
}

func (l *Log) Emerg(v ...interface{}) {
	l.emerg.Print(v...)
}
