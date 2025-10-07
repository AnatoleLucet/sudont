package cmd

import (
	"fmt"
	"os"

	"github.com/AnatoleLucet/sudont/container"
	"github.com/AnatoleLucet/sudont/container/process"
	"github.com/urfave/cli/v3"
)

type runParams struct {
	User string
	Args []string
}

func run(params runParams) error {
	io := process.IO{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	c, err := container.New(container.Config{
		User: params.User,
	})
	if err != nil {
		return fmt.Errorf("unable to create container: %w", err)
	}

	p, err := c.Run(params.Args, io)
	if err != nil {
		return fmt.Errorf("unable to run command: %w", err)
	}

	code, err := p.Wait()
	if err != nil {
		return fmt.Errorf("unable to wait for command: %w", err)
	}

	return cli.Exit("", code)
}
