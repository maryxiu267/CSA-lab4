package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	readLine := bufio.NewReader(conn)
	for { //handles multiple messages from one client.
		msg, _ := readLine.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(conn, "ok")
	}

}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	for { //continuously waits for new clients to connect
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}

}
