package app

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClient(uri string) {
	fmt.Println("Running as client")
	conn, err := net.Dial("tcp", uri) // Replace with listener's address
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Start a goroutine to send data from stdin to server
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text)
		}
	}()

	// Start a goroutine to receive data from server and print to stdout
	go func() {
		reader := bufio.NewReader(conn)
		for {
			text, _ := reader.ReadString('\n')
			fmt.Print(text)
		}
	}()

	// Wait indefinitely
	select {}
}
