/*
Copyright © 2022 Juan Sierra eliotandelon@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "IpTracker",
	Short: "Ip Tracker",
	Long:  "Ip Tracker CLI App",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
