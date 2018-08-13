package main

import (
	"os"

	"github.com/philkra/springboot2beat/cmd"

	_ "github.com/philkra/springboot2beat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
