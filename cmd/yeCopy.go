package cmd

import (
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy krega file ko",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("cp", args...)
	},
}

func init() {
	yeCmd.AddCommand(copyCmd)
}
