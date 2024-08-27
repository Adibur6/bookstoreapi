package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bookapi",
	Short: "BookAPI is a CLI tool to manage your book API server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	// Add any persistent flags or global flags here

}
