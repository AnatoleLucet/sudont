package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func Execute() error {

	flags := pflag.NewFlagSet("sudont", pflag.ContinueOnError)

	user := flags.StringP("user", "u", "", "The user to run the command as")
	version := flags.BoolP("version", "v", false, "Show version")
	help := flags.BoolP("help", "h", false, "Show help")

	flags.Parse(os.Args[1:])

	remainArgs := flags.Args()

	fmt.Println("flags:", *user, *version, *help)
	fmt.Println("args:", remainArgs)

	return nil
}
