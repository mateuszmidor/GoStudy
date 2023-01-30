package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

var rootCmd = &cobra.Command{
	Version:           VERSION,
	Use:               "cobra-cli",
	Short:             "cobra-cli - demo of cobra CLI",
	Long:              "cobra-cli - demo of cobra CLI",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("*** COBRA CLI tool v%s ***\n", VERSION)
		fmt.Println("Use -h for help")
	},
}

func RootExecute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
