package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Disconnected from server")
			return
		}
		fmt.Print(message)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	for { //for loop so u can send multiple messages
		fmt.Print("> ")
		message, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		fmt.Fprintln(conn, message)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	//TODO Try to connect to the server
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Printf("Failed to connect to server at %s: %v\n", *addrPtr, err)
		return
	}

	//TODO Start asynchronously reading and displaying messages
	go read(conn)

	//TODO Start getting and sending user messages.
	write(conn)
}
