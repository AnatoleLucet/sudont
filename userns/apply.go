package userns

import (
	"fmt"
	"os"

	"github.com/AnatoleLucet/sudont/user"
	"golang.org/x/sys/unix"
)

func applyUserNS(ns *UserNS) error {
	// Unset envs so they can be substituted with the new user info.
	os.Unsetenv("HOME")
	os.Unsetenv("USER")
	os.Unsetenv("UID")
	os.Unsetenv("GID")
	os.Unsetenv("SHELL")

	u, err := user.LookupUID(ns.UID)
	if err != nil {
		return err
	}

	// Only set supplementary groups if we're root, as non-root users
	// are not allowed to set supplementary groups.
	if user.IsRoot(unix.Geteuid()) {
		if err := unix.Setgroups(ns.SGIDs); err != nil {
			return fmt.Errorf("unable to set supplementary groups: %w", err)
		}
	}

	if err := unix.Setgid(ns.GID); err != nil {
		return fmt.Errorf("unable to set gid: %w", err)
	}
	if err := unix.Setuid(ns.UID); err != nil {
		return fmt.Errorf("unable to set uid: %w", err)
	}

	if envHome := os.Getenv("HOME"); envHome == "" {
		if err := os.Setenv("HOME", u.Home); err != nil {
			return err
		}
	}

	if envUser := os.Getenv("USER"); envUser == "" {
		if err := os.Setenv("USER", u.Name); err != nil {
			return err
		}
	}

	if envUid := os.Getenv("UID"); envUid == "" {
		if err := os.Setenv("UID", string(rune(u.UID))); err != nil {
			return err
		}
	}

	if envGid := os.Getenv("GID"); envGid == "" {
		if err := os.Setenv("GID", string(rune(u.GID))); err != nil {
			return err
		}
	}

	if envShell := os.Getenv("SHELL"); envShell == "" {
		if err := os.Setenv("SHELL", u.Shell); err != nil {
			return err
		}
	}

	return nil
}
