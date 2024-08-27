package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"os/exec"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running book API server",
	Run: func(cmd *cobra.Command, args []string) {

		stopServer()

	},
}

func init() {

	rootCmd.AddCommand(stopCmd)
}

func stopServer() {

	fmt.Println("Stopping server.")
	cmd := exec.Command("bash", "-c", "ps aux | grep './bookapi' | grep -v 'grep' | awk '{print $2}' | head -n 1")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting the process ID: %v\n", err)
		return
	}

	// Convert output to string and trim any whitespace/newlines
	pid := strings.TrimSpace(string(output))
	if pid == "" {
		fmt.Println("No process found.")
		return
	}

	// Print the process ID
	fmt.Printf("Process ID: %s\n", pid)

	// Step 2: Kill the process
	killCmd := exec.Command("kill", pid)
	err = killCmd.Run()
	if err != nil {
		fmt.Printf("Error killing the process: %v\n", err)
	} else {
		fmt.Println("Process killed successfully.")
	}
}
