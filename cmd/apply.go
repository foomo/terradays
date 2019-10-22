package cmd

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/davecgh/go-spew/spew"
	"github.com/foomo/terradays/terradays"
	"github.com/spf13/cobra"
)

var (
	flagApplyArgs string
	flagTerraform string
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply DIR",
	Short: "builds or changes infrastructure",
	Long:  `Builds or changes infrastructure in sequence according to the Terradays configuration file in DIR.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		terradays.Terraform = flagTerraform

		dirname, err := os.Getwd()
		if err != nil {
			return err
		}
		if len(args) > 0 {
			dirname = args[0]
		}

		plan := &terradays.Plan{}
		if err := plan.Load(dirname); err != nil {
			spew.Dump(err)
			return err
		}

		if err := plan.Apply(dirname, flagApplyArgs); err != nil {
			logrus.WithError(err).Fatal("Could not execute plan")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVar(&flagApplyArgs, "tfargs", "", "Arguements to pass on to Terraform")
	applyCmd.Flags().StringVar(&flagTerraform, "terraform", "terraform", "Path to terraform executable")
}
