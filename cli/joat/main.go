package main

import (
	"fmt"
	"os"

	"github.com/ubiqueworks/joat/cmd"
)

var (
	Build   string
	Version string
)

func main() {
	cmd.Build = Build
	cmd.Version = Version

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
