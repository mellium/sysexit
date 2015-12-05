package honey

import (
	"log/syslog"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"

	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
)

type log struct {
	Files   []logfile  `toml:"file"`
	Console logconsole `toml:"console"`
	Syslog  logsyslog  `toml:"syslog"`
}

type logsyslog struct {
	Level   string `toml:"level"`
	Network string `toml:"network"`
	Raddr   string `toml:"raddr"`
}

type logfile struct {
	Level    string `toml:"level"`
	Filename string `toml:"filename"`
}

type logconsole struct {
	Level string `toml:"level"`
}

// SetupLogging configures the given logger with the supplied log config.
func SetupLogging(logtag string, c log, log *logrus.Logger) {
	// Setup console logging
	if c.Console.Level == "" {
		log.Out = nil
	} else {
		log.Out = os.Stdout
		level, err := logrus.ParseLevel(c.Console.Level)
		if err == nil {
			log.Level = level
		} else {
			log.WithField("level", c.Console.Level).Warn(
				"Unable to parse log level. Falling back to info.")
			log.Level = logrus.InfoLevel
		}
	}

	// Setup log files
	if len(c.Files) > 0 {
		paths := lfshook.PathMap{}
		for _, lf := range c.Files {
			l := log.WithFields(logrus.Fields{
				"file":  lf.Filename,
				"level": lf.Level,
			})
			l.Info("Instantiating log file.")

			level, err := logrus.ParseLevel(lf.Level)
			if err != nil {
				l.Warn("Error while parsing log level. Skipping log file.")
				continue
			}

			paths[level] = lf.Filename
		}
		log.Hooks.Add(lfshook.NewHook(paths))
		log.Formatter = &logrus.TextFormatter{DisableColors: true}
	}

	if c.Syslog.Level != "" {
		l := log.WithFields(logrus.Fields{
			"level":   c.Syslog.Level,
			"network": c.Syslog.Network,
			"raddr":   c.Syslog.Raddr,
		})
		if level, err := logrus.ParseLevel(c.Syslog.Level); err != nil {
			l.Warn("Error while parsing log level. Skipping syslog.")
		} else {
			var priority syslog.Priority
			switch level {
			case logrus.PanicLevel:
				priority = syslog.LOG_ALERT | syslog.LOG_USER
			case logrus.FatalLevel:
				priority = syslog.LOG_CRIT | syslog.LOG_USER
			case logrus.ErrorLevel:
				priority = syslog.LOG_ERR | syslog.LOG_USER
			case logrus.WarnLevel:
				priority = syslog.LOG_WARNING | syslog.LOG_USER
			case logrus.InfoLevel:
				priority = syslog.LOG_INFO | syslog.LOG_USER
			case logrus.DebugLevel:
				priority = syslog.LOG_DEBUG | syslog.LOG_USER
			default:
				l.Panic("Impossible level.")
			}
			hook, err := logrus_syslog.NewSyslogHook(c.Syslog.Network, c.Syslog.Raddr,
				priority, logtag)
			if err != nil {
				l.Warn("Unable to connect to syslog daemon")
			} else {
				log.Hooks.Add(hook)
				log.Formatter = &logrus.TextFormatter{DisableColors: true}
			}
		}
	}
}
