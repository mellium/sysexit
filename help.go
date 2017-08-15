// Copyright 2016 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"mellium.im/cli"
)

func helpCmd(cs *cli.CommandSet) *cli.Command {
	return &cli.Command{
		Usage:       "help [command]",
		Description: `Print detailed usage information about subcommands.`,
		Run: func(c *cli.Command, args ...string) error {
			// If the only argument is the command name (or there aren't any, which
			// shouldn't happen), print the main command help.
			if len(args) < 1 {
				cs.Help(os.Stdout)
				return nil
			}

			// Print the help for the provided subcommand or help topic.
			for _, cmd := range cs.Commands {
				if cmd.Name() != args[0] {
					continue
				}
				cmd.Help(os.Stdout)
				return nil
			}
			return fmt.Errorf("No such help topic %s\n", args[0])
		},
	}
}
