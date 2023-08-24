package commands

import (
	"github.com/spf13/cobra"
)

func init() {

}

// RootCMD is the root command for ctm
var RootCMD = &cobra.Command{
	Use:   "tasks",
	Short: "Ctm is a CLI for managing your TODOs.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
