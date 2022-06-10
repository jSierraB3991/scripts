package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Create files of configuration",
	Long: "Create files of configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cfg called")
	},
}

func init() {
	mkCmd.AddCommand(cfgCmd)
}
