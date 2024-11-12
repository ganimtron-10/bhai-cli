package cmd

import (
	"github.com/spf13/cobra"
)

var banaCmd = &cobra.Command{
	Use:   "bana",
	Short: "File bana denga",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("touch", args...)
	},
}

func init() {
	yeCmd.AddCommand(banaCmd)
}
