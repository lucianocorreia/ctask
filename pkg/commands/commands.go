package commands

import (
	"github.com/lucianocorreia/ctm/pkg/db"
	"github.com/lucianocorreia/ctm/pkg/utils"
	"github.com/spf13/cobra"
)

func init() {
	// flags
	addCMD.Flags().StringP("project", "p", "", "Project name")

	RootCMD.AddCommand(addCMD)
}

// RootCMD is the root command for ctm
var RootCMD = &cobra.Command{
	Use:   "tcm",
	Short: "Ctm is a CLI for managing your TODOs.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var addCMD = &cobra.Command{
	Use:   "add NAME",
	Short: "Add a new task with an optional project name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := db.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}
		defer t.DB.Close()

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}

		if err := t.Insert(args[0], project); err != nil {
			return err
		}

		return nil
	},
}
