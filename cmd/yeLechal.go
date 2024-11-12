package cmd

import (
	"github.com/spf13/cobra"
)

var lechalCmd = &cobra.Command{
	Use:   "lechal",
	Short: "Move krega file ko",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("mv", args...)
	},
}

func init() {
	yeCmd.AddCommand(lechalCmd)
}
