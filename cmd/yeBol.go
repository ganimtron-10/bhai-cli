package cmd

import (
	"github.com/spf13/cobra"
)

var bolCmd = &cobra.Command{
	Use:   "bol",
	Short: "Bol ke dikahyega",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("echo", args...)
	},
}

func init() {
	yeCmd.AddCommand(bolCmd)
}
