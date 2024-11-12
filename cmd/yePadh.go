package cmd

import (
	"github.com/spf13/cobra"
)

var padhCmd = &cobra.Command{
	Use:   "padh",
	Short: "File Padh ke Dikhayega",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("cat", args...)
	},
}

func init() {
	yeCmd.AddCommand(padhCmd)
}
