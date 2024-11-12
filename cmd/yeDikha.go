package cmd

import (
	"github.com/spf13/cobra"
)

var sab bool = false

var dikhaCmd = &cobra.Command{
	Use:   "dikha",
	Short: "Files aur directories dikhayega",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if sab {
			args = append(args, "-al")
		}
		runCommand("ls", args...)
	},
}

func init() {
	yeCmd.AddCommand(dikhaCmd)

	dikhaCmd.Flags().BoolVar(&sab, "sab", false, "Sare Suchna dikahyega")
}
