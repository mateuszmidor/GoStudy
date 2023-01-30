package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// cli command flags
var verbose bool
var outputBinary bool

var addCmd = &cobra.Command{
	Use:     "add <arg1> <arg2>",
	Short:   "Add two numbers",
	Example: "add 2 3",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		arg1, _ := strconv.ParseInt(args[0], 10, 32)
		arg2, _ := strconv.ParseInt(args[1], 10, 32)
		result := arg1 + arg2
		if verbose {
			fmt.Printf("Adding: %d + %d\n", arg1, arg2)
		}
		if outputBinary {
			fmt.Printf("%b\n", result)
		} else {
			fmt.Printf("%d\n", result)
		}
	},
}

var mulCmd = &cobra.Command{
	Use:     "mul <arg1> <arg2>",
	Short:   "Multiply two numbers",
	Example: "mul 2 3",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		arg1, _ := strconv.ParseInt(args[0], 10, 32)
		arg2, _ := strconv.ParseInt(args[1], 10, 32)
		result := arg1 * arg2
		if verbose {
			fmt.Printf("Multiplying: %d * %d\n", arg1, arg2)
		}
		if outputBinary {
			fmt.Printf("%b\n", result)
		} else {
			fmt.Printf("%d\n", result)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")

	// add command
	addCmd.Flags().BoolVar(&outputBinary, "bin", false, "print output in binary format")
	rootCmd.AddCommand(addCmd)

	// mul command
	mulCmd.Flags().BoolVar(&outputBinary, "bin", false, "print output in binary format")
	rootCmd.AddCommand(mulCmd)

}
