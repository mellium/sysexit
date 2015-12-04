package main

import (
	"bitbucket.org/SamWhited/honey/config"

	"fmt"
)

func main() {

	// Load the config

	c := config.Default()

	err := c.LoadFile("config.toml")
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err.Error())
	}

	fmt.Printf("%#v\n", c)
}
