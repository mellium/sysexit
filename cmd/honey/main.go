package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"

	"bitbucket.org/SamWhited/honey"
)

var configFile string

func init() {
	// Setup flags
	flag.StringVar(&configFile, "config", "config.toml", "Selects the config file")
	flag.Parse()
}

func main() {

	// Load the config
	c, err := honey.ConfigFromFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Warning:", err)
		} else {
			fmt.Println("Error while loading config: ", err)
			os.Exit(1)
		}
	}

	// Setup logging
	log := honey.SetupLogging("honey", c.Log)

	for _, v := range c.VHosts {
		server := honey.ServerFromConfig(v)
		err := server.ListenAndServe()
		if err != nil {
			log.Error(err)
		}
	}

	select {}
}
