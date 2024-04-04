package cmd

import (
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "mem",
	Short: "mem",
	Long:  "resource testing tools - memory",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(memoryCmd)
}
