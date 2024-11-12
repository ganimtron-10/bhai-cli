package cmd

import (
	"github.com/spf13/cobra"
)

var yeCmd = &cobra.Command{
	Use:   "ye",
	Short: "Bhai ko kaam krne mein madat kre",
	// Long: ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	bhaiCmd.AddCommand(yeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
