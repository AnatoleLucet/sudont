package container

import (
	"fmt"
	"os/exec"

	"github.com/AnatoleLucet/sudont/container/process"
	"github.com/AnatoleLucet/sudont/user"
)

type Container struct {
	User *user.User
	Env  []string
}

type Config struct {
	User string
	Env  []string
}

func New(config Config) (*Container, error) {
	u, err := user.Lookup(config.User)
	if err != nil {
		return nil, err
	}

	return &Container{
		User: u,
		Env:  config.Env,
	}, nil
}

func (c *Container) Run(args []string, io process.IO) (*process.Process, error) {
	p, err := process.New(process.ProcessOpts{
		IO: io,

		Args: args,
		Env:  c.Env,

		UID:   c.User.UID,
		GID:   c.User.GID,
		SGIDs: c.User.SGIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to create process: %w", err)
	}

	if err := c.Spawn(p); err != nil {
		return nil, fmt.Errorf("unable to spawn process: %w", err)
	}

	return p, nil
}

func (c *Container) Spawn(proc *process.Process) error {
	if err := proc.UserNS.Apply(); err != nil {
		return fmt.Errorf("unable to apply user namespace: %w", err)
	}

	bin, err := exec.LookPath(proc.Args[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(bin, proc.Args[1:]...)
	cmd.Env = proc.Env
	cmd.Stdin = proc.IO.Stdin
	cmd.Stdout = proc.IO.Stdout
	cmd.Stderr = proc.IO.Stderr
	if err := cmd.Start(); err != nil {
		return err
	}

	proc.PID = cmd.Process.Pid

	// TODO: check if we can restore the user namespace after spawning the process
	// if err := proc.UserNS.Restore(); err != nil {
	// 	return fmt.Errorf("unable to restore user namespace: %w", err)
	// }

	return nil
}
