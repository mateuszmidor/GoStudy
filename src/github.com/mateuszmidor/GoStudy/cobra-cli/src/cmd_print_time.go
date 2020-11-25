package main

import (
	"time"

	"github.com/spf13/cobra"
)

func cmdPrintTime() *cobra.Command {
	return &cobra.Command{
		Use: "curtime", // command name to be provided in shell, eg "> cobra-cli curtime"
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("The time is ", prettyTime)
			return nil
		},
	}
}
