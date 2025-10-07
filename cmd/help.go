package cmd

import "runtime"

const VERSION = "1.0.0-rc.1"

func version() string {
	return VERSION + " (" + runtime.Version() + " on " + runtime.GOOS + "/" + runtime.GOARCH + "; " + runtime.Compiler + ")"
}

func help() string {
	v := version()
	h := `sudont - Making sure a command is never run as root.

Usage:
  [global options] <command> [command arguments...]

Version:
  ` + v + `

Global Options:
  --user string, -u string  The user to run the command as
  --version, -v             Show version
  --help, -h                Show help
`

	return h
}
