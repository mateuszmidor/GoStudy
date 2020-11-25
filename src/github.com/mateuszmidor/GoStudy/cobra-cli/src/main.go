package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// cli will dispatch commands
	cli := &cobra.Command{
		Use:          "cobra-cli",                                 // this should be the name of the program
		Short:        "This is 'cobra' package demo for CLI apps", // description of the cli program
		SilenceUsage: true,                                        // dont print usage when error occurs
	}

	// add command to cli
	cli.AddCommand(cmdPrintTime())

	// handle command or printout error/help if wrong command/no command specified
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
