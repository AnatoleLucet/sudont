package process

import (
	"io"
	"os"

	"github.com/AnatoleLucet/sudont/userns"
)

type IO struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type Process struct {
	IO

	PID int

	Args []string
	Env  []string

	UserNS userns.UserNS
}

type ProcessOpts struct {
	IO

	Args []string
	Env  []string

	UID   int
	GID   int
	SGIDs []int
}

func New(opts ProcessOpts) (*Process, error) {
	ns, err := userns.New(opts.UID, opts.GID, opts.SGIDs)
	if err != nil {
		return nil, err
	}

	return &Process{
		IO:     opts.IO,
		Args:   opts.Args,
		Env:    opts.Env,
		UserNS: *ns,
	}, nil
}

func (p *Process) Wait() (code int, err error) {
	if p.PID == 0 {
		return 0, nil
	}

	ps, err := os.FindProcess(p.PID)
	if err != nil {
		return 0, err
	}

	state, err := ps.Wait()
	if err != nil {
		return 0, err
	}

	return state.ExitCode(), nil
}
