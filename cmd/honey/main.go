package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	flag "github.com/ogier/pflag"

	"bitbucket.org/SamWhited/honey"
)

var configFile string
var log *logrus.Logger

func init() {
	// Setup default logging options
	log = logrus.New()

	// Setup flags
	flag.StringVar(&configFile, "config", "config.toml", "Selects the config file")
	flag.Parse()
}

func main() {

	// Load the config
	c, err := honey.ConfigFromFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info(err)
		} else {
			log.Fatal("Error while loading config: ", err)
		}
	}

	honey.SetupLogging("honey", c.Log, log)
}
