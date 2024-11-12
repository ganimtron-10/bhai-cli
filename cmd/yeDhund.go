package cmd

import (
	"github.com/spf13/cobra"
)

var dhundCmd = &cobra.Command{
	Use:   "dhund",
	Short: "File and Directory mein talaash kre",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		runCommand("find", args...)
	},
}

func init() {
	yeCmd.AddCommand(dhundCmd)
}
