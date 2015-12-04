package logging

import (
	"log/syslog"
	"os"
)

// An Option is used to set the behavior of a logger.
type Option func(o *options)

type options struct {
	priority syslog.Priority
	filename string
	filemode os.FileMode
	stdout   bool
	stderr   bool
	syslog   bool
	raddr    string
	network  string
}

func getOpts(o ...Option) (res options) {
	for _, f := range o {
		f(&res)
	}
	if res.filemode == 0 {
		FileMode(0640)(&res)
	}
	return
}

// Priority sets the given syslog priority (a combination of facility and
// severity) and enables syslog.
func Priority(p syslog.Priority) Option {
	return func(o *options) {
		o.syslog = true
		o.priority = p
	}
}

// Filename sets the log file to the given path.
func Filename(f string) Option {
	return func(o *options) {
		o.filename = f
	}
}

// Network sets the network for a remote syslog server.
func Network(network string) Option {
	return func(o *options) {
		o.syslog = true
		o.network = network
	}
}

// RAddr sets the relay address to use for syslog connections.
func RAddr(raddr string) Option {
	return func(o *options) {
		o.syslog = true
		o.raddr = raddr
	}
}

// FileMode sets the mode (0640 if no option is specified) for the log file.
func FileMode(mode os.FileMode) Option {
	return func(o *options) {
		o.filemode = mode
	}
}

var (
	// Stdout enables logging output to STDOUT.
	Stdout Option = func(o *options) {
		o.stdout = true
	}
	// Stderr enables logging output to STDOUT.
	Stderr Option = func(o *options) {
		o.stderr = true
	}
)
