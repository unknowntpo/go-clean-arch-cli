package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get the task with id or name specified",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
