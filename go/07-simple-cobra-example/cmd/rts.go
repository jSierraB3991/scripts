package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rtsCmd = &cobra.Command{
	Use:   "rts",
	Short: "Create files of routes",
	Long: "Create files of routes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rts called")
	},
}

func init() {
	mkCmd.AddCommand(rtsCmd)
}
