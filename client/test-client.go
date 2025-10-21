package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	fmt.Printf(message)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")

	for { //for loop so u can send multiple messages`
		fmt.Print("> ")
		message, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, message)
		readConn(conn) //waits to receive something back from the server
	}

}
