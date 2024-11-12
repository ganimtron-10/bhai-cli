package cmd

import (
	"github.com/spf13/cobra"
)

var saafCmd = &cobra.Command{
	Use:   "saaf",
	Short: "Terminal Saaf kre",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("clear", args...)
	},
}

func init() {
	yeCmd.AddCommand(saafCmd)
}
