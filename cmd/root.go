package cmd

import (
	"github.com/christopherime/go-project-generator/internal/ui"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopg",
	Short: "Golang project generator",
	Run: func(cmd *cobra.Command, args []string) {
		ui.StartPrompt()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
