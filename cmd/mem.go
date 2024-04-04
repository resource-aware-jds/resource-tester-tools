/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// memCmd represents the mem command
var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "mem to reserve",
	Long:  `memory to researve"`,
	Run: func(cmd *cobra.Command, args []string) {
		MBToReserve, _ := cmd.Flags().GetInt("mb")
		a := make([]byte, MBToReserve*1000000)
		fmt.Println("Reserved: ", " : ", MBToReserve, "mb")
		fmt.Println("Press Ctrl+C to abort this operation")

		for i := range a {
			a[i] = 0x99
		}

		// Gracefully Shutdown
		// Make channel listen for signals from OS
		gracefulStop := make(chan os.Signal, 1)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		<-gracefulStop
	},
}

func init() {
	rootCmd.AddCommand(memCmd)

	memCmd.Flags().Int("mb", 100, "Memory to reserve")
}
