package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Client chan<- string

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string)
)

var (
	host = flag.String("h", "localhost", "hostname")
	port = flag.String("p", "3090", "port")
)

// Client1 -> Server -> HanldeConnection(Client1)
// HanldeConnnection handles the client connection of a single user
func HanldeConnection(conn net.Conn) {
	defer conn.Close()

	// Create the client channel
	message := make(chan string)
	go MessageWrite(conn, message)

	// Get the client's name
	clientName := conn.RemoteAddr().String()

	// Send just to the client his name
	message <- "Welcome to the server, your name is " + clientName

	// Send the message to all the clients
	messages <- clientName + " has joined!"

	// Add the client to the list of clients
	entering <- message

	// Read the messages from the client
	// If the loop breaks, that means that the client has disconnected
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- clientName + ": " + inputMessage.Text()
	}

	// Remove the client from the list of clients
	leaving <- message
	messages <- clientName + " has left!"
}

// MessageWrite recives messages from the channel and writes them to the client
func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

// Broadcast sends the message to all the clients, and handles incoming
// and outgoing connections
func Broadcast() {

	// Map the clients to a boolean
	clients := make(map[Client]bool)

	for {
		// Multiplex the messages
		select {
		// We get a new message
		case msg := <-messages:
			// Send the message to all the clients
			for client := range clients {
				client <- msg
			}
		// We get a new client
		case client := <-entering:
			clients[client] = true
		// A client that has disconnected
		case client := <-leaving:
			delete(clients, client)
			close(client)
		}
	}
}

func main() {
	flag.Parse()

	// Create the server and listen to it
	listener, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// Start the broadcast
	go Broadcast()

	// Listen for connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go HanldeConnection(conn)
	}
}
