package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Create files of libraries",
	Long: "Create files of libraries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lib called")
	},
}

func init() {
	mkCmd.AddCommand(libCmd)
}
