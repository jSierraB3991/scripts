package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ctlCmd = &cobra.Command{
	Use:   "ctl",
	Short: "Create files of controllers",
	Long: "Create files of controllers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ctl called")
	},
}

func init() {
	mkCmd.AddCommand(ctlCmd)
}
