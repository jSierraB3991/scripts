package cmd

import (
	"github.com/OkabeRitarou/07-cli-javic-cobra/service"
	"github.com/spf13/cobra"
)

var mdlCmd = &cobra.Command{
	Use:   "mdl",
	Short: "Create files of models",
	Long:  "Create files of models",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("crud"); flags {
			service.VarSrv.CreateMdlCrud(args[0])
		} else {
			service.VarSrv.CreateMdlDefault(args[0])
		}
	},
}

func init() {
	mdlCmd.Flags().BoolP("crud", "c", false, "create crud template")
	mkCmd.AddCommand(mdlCmd)
}
