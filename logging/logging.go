package logging

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

// New creates a new Logger that will output to several locations at once
// (eg. to stdout, a file, and syslog). Flags can be found in the log package.
func NewLogger(prefix string, flag int, options ...Option) (*log.Logger, error) {
	o := getOpts(options...)
	writers := []io.Writer{}

	if o.stdout {
		writers = append(writers, os.Stdout)
	}

	if o.stderr {
		writers = append(writers, os.Stderr)
	}

	if o.filename != "" {
		f, err := os.OpenFile(o.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, o.filemode)
		if err != nil {
			return nil, err
		}
		writers = append(writers, f)
	}

	if o.syslog {
		sw, err := syslog.Dial(o.network, o.raddr, o.priority, prefix)
		if err != nil {
			return nil, err
		}
		writers = append(writers, sw)
	}

	return log.New(io.MultiWriter(writers...), prefix, flag), nil
}
