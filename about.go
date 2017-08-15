// Copyright 2017 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"runtime"

	"mellium.im/cli"
)

func aboutCmd(configFile, version, commit string) *cli.Command {
	return &cli.Command{
		Usage:       "about",
		Description: "Show information about Mellium.",
		Run: func(c *cli.Command, _ ...string) error {
			fmt.Printf(`Mellium (%s)

version:     %s
git hash:    %s
go version:  %s
go compiler: %s
platform:    %s/%s

config file: %s
`,
				os.Args[0],
				version, commit,
				runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
				configFile)
			return nil
		},
	}
}
