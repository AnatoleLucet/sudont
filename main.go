package main

import (
	"fmt"
	"os"

	"github.com/AnatoleLucet/sudont/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(fmt.Errorf("sudont: %w", err))
		os.Exit(1)
	}
}
