// Package app is the main package for the application.
package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/gosh/internal/meta"
	"github.com/urfave/cli/v2"
)

// Run is the entry point for the application.
func Run(args []string) int {
	app := cli.NewApp()
	app.Name = meta.Name
	app.Version = meta.Version
	app.Usage = meta.Description
	app.HideHelpCommand = true

	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:    "path-length",
			Aliases: []string{"l"},
			Usage:   "set the number of trailing path parts to display",
			Value:   3,
			EnvVars: []string{"PROMPT_DIRTRIM"},
		},
		&cli.BoolFlag{
			Name:    "ssh-segment",
			Aliases: []string{"ssh"},
			Usage:   "enable the SSH segment",
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "git-segment",
			Aliases: []string{"git"},
			Usage:   "enable the git segment",
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "privilege-segment",
			Aliases: []string{"priv"},
			Usage:   "enable the privilege segment",
			Value:   true,
		},
	}

	app.Action = PromptAction

	if err := app.Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		return 1
	}

	return 0
}
