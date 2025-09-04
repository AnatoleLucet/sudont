package main

import (
	"fmt"
	"os"

	"github.com/AnatoleLucet/sudont/cmd"
	"github.com/AnatoleLucet/sudont/container"
)

func init() {
	container.Init()
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(fmt.Errorf("sudont: %w", err))
		os.Exit(1)
	}
}
