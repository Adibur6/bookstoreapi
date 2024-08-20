package cmd

import (
	"github.com/adibur6/bookstoreapi/config"
	"github.com/spf13/cobra"
	"os"
)

var Cfg config.Config
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
	Cfg = *config.LoadConfig(config.Configfile)
	// Add any persistent flags or global flags here

}
