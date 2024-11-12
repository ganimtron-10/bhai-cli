package cmd

import (
	"github.com/spf13/cobra"
)

var nikalCmd = &cobra.Command{
	Use:   "nikal",
	Short: "File aur Directory nikal denga",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("rm", args...)
	},
}

func init() {
	yeCmd.AddCommand(nikalCmd)
}
