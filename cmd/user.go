package cmd

import (
	"errors"
	"os"
	"os/user"
)

func findUserFromSudo() string {
	return os.Getenv("SUDO_USER")
}

func findUserFromNamespace() string {
	u, _ := user.Current()
	return u.Name
}

func findNonRootUser(name string) (string, error) {
	if name == "" || name == "root" {
		name = findUserFromSudo()
	}
	if name == "" || name == "root" {
		name = findUserFromNamespace()
	}

	if name == "root" {
		return "", errors.New("refusing to run as root user")
	}

	return name, nil
}
