package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	flag "github.com/ogier/pflag"

	"bitbucket.org/SamWhited/honey/config"
)

var configFile string

func init() {

	// Setup default logging options
	log.SetLevel(log.InfoLevel)

	// Setup flags
	flag.StringVar(&configFile, "config", "config.toml", "Selects the config file")
	flag.Parse()
}

func main() {

	// Load the config
	err = c.LoadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info(err)
		} else {
			log.Fatal("Error while loading config: ", err)
		}
	}

	log.Infof("%+v\n", c)
}
