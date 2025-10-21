package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

type Message struct {
	sender  int
	message string
}

// handleError now stops execution for fatal errors
func handleErrorFatal(err error, context string) {
	if err != nil {
		fmt.Printf("%s: %v\n", context, err)
		panic(err) // stop execution if fatal
	}
}

// handleClientError handles client disconnects
func handleClientError(err error, clientID int) bool {
	if err != nil {
		fmt.Printf("Client %d disconnected\n", clientID)
		return true // signal to stop
	}
	return false
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		conns <- conn
	}
}

func handleClient(client net.Conn, clientID int, msgs chan Message) {
	reader := bufio.NewReader(client)
	for {
		msg, err := reader.ReadString('\n')
		if handleClientError(err, clientID) {
			client.Close()
			return
		}
		msg = strings.TrimRight(msg, "\n")
		msgs <- Message{clientID, msg}
	}
}

func main() {
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	ln, err := net.Listen("tcp", *portPtr)
	handleErrorFatal(err, "Failed to start listener")

	conns := make(chan net.Conn)
	msgs := make(chan Message)
	clients := make(map[int]net.Conn)

	go acceptConns(ln, conns)

	var nextID int
	for {
		select {
		case conn := <-conns:
			id := nextID
			nextID++
			clients[id] = conn
			go handleClient(clients[id], id, msgs)

		case msg := <-msgs:
			for id, conn := range clients {
				if msg.sender != id {
					_, err := fmt.Fprintln(conn, msg.message)
					if err != nil {
						fmt.Printf("Failed to send message to client %d: %v\n", id, err)
						conn.Close()
						delete(clients, id)
					}
				}
			}
		}
	}
}
