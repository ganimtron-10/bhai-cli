package cmd

import (
	"github.com/spf13/cobra"
)

var kahaCmd = &cobra.Command{
	Use:   "kaha",
	Short: "Abhi kaha ho wo batayega",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		runCommand("pwd", args...)
	},
}

func init() {
	yeCmd.AddCommand(kahaCmd)
}
