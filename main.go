// Package main is just an entry point.
package main

import (
	"fmt"
	"os"

	"github.com/SidhuG/hellogopher/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
