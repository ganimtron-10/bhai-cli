package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var bhaiCmd = &cobra.Command{
	Use:   "bhai",
	Short: "BhaiCLI - CLI ka bhai, Bhai ka CLI",
	Long: `
$$$$$$$\  $$\   $$\  $$$$$$\  $$$$$$\  $$$$$$\  $$\       $$$$$$\ 
$$  __$$\ $$ |  $$ |$$  __$$\ \_$$  _|$$  __$$\ $$ |      \_$$  _|
$$ |  $$ |$$ |  $$ |$$ /  $$ |  $$ |  $$ /  \__|$$ |        $$ |  
$$$$$$$\ |$$$$$$$$ |$$$$$$$$ |  $$ |  $$ |      $$ |        $$ |  
$$  __$$\ $$  __$$ |$$  __$$ |  $$ |  $$ |      $$ |        $$ |  
$$ |  $$ |$$ |  $$ |$$ |  $$ |  $$ |  $$ |  $$\ $$ |        $$ |  
$$$$$$$  |$$ |  $$ |$$ |  $$ |$$$$$$\ \$$$$$$  |$$$$$$$$\ $$$$$$\ 
\_______/ \__|  \__|\__|  \__|\______| \______/ \________|\______|

`,
}

func Execute() {
	err := bhaiCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
