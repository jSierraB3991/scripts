package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mdlCmd = &cobra.Command{
    Use:   "mdl",
    Short: "Create files of models",
    Long: "Create files of models",
    Run: func(cmd *cobra.Command, args []string) {
        if flags, _ := cmd.Flags().GetBool("crud"); flags {
            fmt.Println("crud activado")
        }
        fmt.Println("mdl called")
    },
}

func init() {
    mdlCmd.Flags().BoolP("crud", "c", false, "create crud template")
    mkCmd.AddCommand(mdlCmd)
}
