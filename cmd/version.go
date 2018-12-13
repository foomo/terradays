package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints out version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(rootCmd.Use + " " + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
