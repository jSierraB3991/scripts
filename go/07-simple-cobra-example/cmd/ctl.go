package cmd

import (
	"github.com/OkabeRitarou/07-cli-javic-cobra/service"
	"github.com/spf13/cobra"
)

var ctlCmd = &cobra.Command{
	Use:   "ctl",
	Short: "Create files of controllers",
	Long:  "Create files of controllers",
	Run: func(cmd *cobra.Command, args []string) {
		if flags, _ := cmd.Flags().GetBool("crud"); flags {
			service.VarSrv.CreateCtlCrud(args[0])
		} else {
			service.VarSrv.CreateCtlDefault(args[0])
		}
	},
}

func init() {
	mkCmd.AddCommand(ctlCmd)
}
