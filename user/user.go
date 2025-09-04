package user

import (
	"fmt"
	"strconv"

	"github.com/moby/sys/user"
	"golang.org/x/sys/unix"
)

type User struct {
	Name string

	UID   int
	GID   int
	SGIDs []int

	Home  string
	Shell string
}

// Wrapper around github.com/moby/sys/user.GetExecUser
func Lookup(name string) (*User, error) {
	passwdPath, err := user.GetPasswdPath()
	if err != nil {
		return nil, err
	}
	groupPath, err := user.GetGroupPath()
	if err != nil {
		return nil, err
	}

	execUser, err := user.GetExecUserPath(name, &user.ExecUser{}, passwdPath, groupPath)
	if err != nil {
		return nil, fmt.Errorf("unable to find user %q: %w", name, err)
	}

	passwdUser, err := user.LookupUid(execUser.Uid)
	if err != nil {
		return nil, fmt.Errorf("unable to find user %q in passwd file: %w", name, err)
	}

	return &User{
		Name: passwdUser.Name,

		UID:   execUser.Uid,
		GID:   execUser.Gid,
		SGIDs: execUser.Sgids,

		Home:  passwdUser.Home,
		Shell: passwdUser.Shell,
	}, nil
}

func LookupUID(uid int) (*User, error) {
	return Lookup(strconv.Itoa(uid))
}

func Current() (*User, error) {
	return LookupUID(unix.Getuid())
}

func IsRoot(uid int) bool {
	return uid == 0
}
