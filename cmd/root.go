package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcac",
	Short: "go-clean-arch-cli is a cli tool for go-clean-arch",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gcac is called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
