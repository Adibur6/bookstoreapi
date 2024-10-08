package cmd

import (
	"fmt"
	"github.com/adibur6/bookstoreapi/apihandler"
	"github.com/adibur6/bookstoreapi/utility"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var port int = 8080
var detached bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the book API server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server at ", port)
		if detached {
			startDetached()
		} else {
			startServer()
		}

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().IntVarP(&port, "port", "p", 8080, "Specify port to run the server on")
	startCmd.Flags().BoolVarP(&detached, "detached", "d", false, "Run server in detached mode")
}

func startServer() {
	if !utility.IsPortAvailable(port) {
		fmt.Printf("Server is already running on port %d.\n", port)
		return
	}

	apihandler.Start(port)

}

func startDetached() {
	if !utility.IsPortAvailable(port) {
		fmt.Printf("Server is already running on port %d.\n", port)
		return
	}

	// Create the command to start the server in detached mode
	cmd := exec.Command(os.Args[0], "start", "-p", strconv.Itoa(port))

	// Redirect stdout and stderr to avoid cluttering the terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	err := cmd.Start()
	if err != nil {
		log.Fatal("Failed to start server in detached mode:", err)
	}

	fmt.Printf("Server started in detached mode on port %d with PID %d\n", port, cmd.Process.Pid)

	// Optionally, you can detach the process if needed
	err = cmd.Process.Release()
	if err != nil {
		log.Fatal("Failed to release process:", err)
	}

}
