package cmd

import (
	"fmt"
	"github.com/adibur6/bookstoreapi/config"
	"github.com/adibur6/bookstoreapi/utility"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running book API server",
	Run: func(cmd *cobra.Command, args []string) {
		Cfg = *config.LoadConfig(config.Configfile)
		stopServer()

	},
}

func init() {

	rootCmd.AddCommand(stopCmd)
}

func stopServer() {
	if utility.IsPortAvailable(Cfg.PortNumber) {
		fmt.Println(Cfg.PortNumber)
		fmt.Println("No server is currently running.")
		return
	}

	fmt.Printf("Stopping server on port %d...\n", Cfg.PortNumber)

	// Find and kill the process using the port number
	cmd := exec.Command("sh", "-c", fmt.Sprintf("kill -9 $(lsof -t -i :%d)", Cfg.PortNumber))
	err := cmd.Run()
	if err != nil {
		log.Fatal("Failed to stop server:", err)
	}

}
