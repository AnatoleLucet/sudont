package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

const NAME = "sudont"
const USAGE = "Making sure a command is never run as root."
const DESCRIPTION = `sudont is a tool to run commands as a non-root user.
It tries to automatically determine a non-root user, and refuses to run commands as root, even if explicitly asked to do so.`

func Execute() error {
	cmd := &cli.Command{
		Name:        NAME,
		Usage:       USAGE,
		Description: DESCRIPTION,
		Version:     version(),
		ArgsUsage:   "<command> [command arguments...]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "Specify the user to run the command as.",
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "command",
				UsageText: "The command to run as a non-root user.",
				Config:    cli.StringConfig{TrimSpace: true},
			},
			&cli.StringArgs{
				Min:       0,
				Max:       -1,
				Name:      "arguments",
				UsageText: "Arguments to pass to the command.",
				Config:    cli.StringConfig{TrimSpace: true},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			command := cmd.StringArg("command")
			if command == "" {
				cli.ShowAppHelp(cmd)
				return errors.New("no command specified")
			}

			name, err := findNonRootUser(cmd.String("user"))
			if err != nil {
				return fmt.Errorf("unable to determine current user. Make sure a non-root user exists and please specify a user with --user. %w", err)
			}

			args := cmd.StringArgs("arguments")

			return run(runOpts{
				User: name,
				Args: append([]string{command}, args...),
			})
		},
	}

	return cmd.Run(context.Background(), os.Args)
}
