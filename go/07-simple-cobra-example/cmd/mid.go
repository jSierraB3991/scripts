package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var midCmd = &cobra.Command{
	Use:   "mid",
	Short: "Create files of middlewares",
	Long: "Create files of middlewares",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mid called")
	},
}

func init() {
	mkCmd.AddCommand(midCmd)
}
