// Copyright 2016 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"

	"mellium.im/cli"
)

// Build configuration variables.
var (
	defConfigFile = "config.toml"
	version       = "devel"
	commit        = ""
)

func main() {
	var (
		configFile string
	)

	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&configFile, "config", defConfigFile, "the config file to load")
	flags.Parse(os.Args[1:])

	logger := log.New(os.Stderr, "mel ", log.LstdFlags)

	// Commands
	cmds := &cli.CommandSet{
		Name:  os.Args[0],
		Flags: flags,
		Commands: []*cli.Command{
			aboutCmd(configFile, version, commit),
		},
	}
	cmds.Commands = append(cmds.Commands, helpCmd(cmds))

	err := cmds.Run(flags.Args()...)
	switch {
	case err == flag.ErrHelp:
		// Ignore errors from using an implicitly defined -h or -help flag.
		return
	case err != nil:
		logger.Fatal(err)
	}
}
