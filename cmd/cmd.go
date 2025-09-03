package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	"github.com/urfave/cli/v3"
)

type runParams struct {
	user *user.User
	bin  string
	args []string
}

func Execute() error {
	cmd := &cli.Command{
		Name:      "sudont",
		Usage:     "When you want to make sure a command is NOT run as root.",
		ArgsUsage: "<command> [command arguments...]",
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
				err := cli.ShowAppHelp(cmd)
				if err != nil {
					return err
				}

				return errors.New("no command specified")
			}

			user, err := findUser(cmd.String("user"))
			if err != nil {
				return err
			}

			args := cmd.StringArgs("arguments")

			return run(runParams{
				user: user,
				bin:  command,
				args: args,
			})
		},
	}

	return cmd.Run(context.Background(), os.Args)
}

func isRoot() bool {
	return os.Geteuid() == 0
}

func findNonRootUser() (*user.User, error) {
	u := os.Getenv("SUDO_USER")
	if u == "" {
		return nil, errors.New("unable to determine non-root user. Please set the SUDO_USER environment variable.")
	}

	return user.Lookup(u)
}

func findUserByName(name string) (*user.User, error) {
	if name == "root" {
		return nil, errors.New("refusing to run as root user")
	}

	u, err := user.Lookup(name)
	if err != nil {
		return nil, fmt.Errorf("unable to find user %q: %w", name, err)
	}

	return u, nil
}

func findUser(name string) (*user.User, error) {
	if !isRoot() {
		return user.Current()
	}

	if name != "" {
		return findUserByName(name)
	}

	return findNonRootUser()
}

func run(params runParams) error {
	uid, err := strconv.Atoi(params.user.Uid)
	if err != nil {
		return fmt.Errorf("invalid uid %q: %w", params.user.Uid, err)
	}

	gid, err := strconv.Atoi(params.user.Gid)
	if err != nil {
		return fmt.Errorf("invalid gid %q: %w", params.user.Gid, err)
	}

	syscall.Setgid(gid)
	syscall.Setuid(uid)

	c := exec.Command(params.bin, params.args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
