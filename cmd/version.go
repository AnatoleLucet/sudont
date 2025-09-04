package cmd

import "runtime"

const VERSION = "1.0.0-rc.1"

func version() string {
	// 1.17 (go1.18.2 on linux/amd64; gc)
	return VERSION + " (" + runtime.Version() + " on " + runtime.GOOS + "/" + runtime.GOARCH + "; " + runtime.Compiler + ")"
}
