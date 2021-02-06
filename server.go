package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	clientCount := 0

	allClients := make(map[net.Conn]string)

	newConnections := make(chan net.Conn)

	deadConnections := make(chan net.Conn)

	messages := make(chan string)

	server, err := net.Listen("tcp", ":6000")

	if err != nil {
		panic(err)
	}

	defer server.Close()

	fmt.Println("Server started on port 6000")

	go func() {
		for {
			conn, err := server.Accept()

			if err != nil {
				panic(err)
			}
			newConnections <- conn
		}
	}()

	for {
		select {
		case conn := <-newConnections:
			log.Println("Waiting nickname for new client")

			go func(conn net.Conn) {
				conn.Write([]byte("Welcome to the club\n"))
				conn.Write([]byte("Choose a nickname :\n"))
				reader := bufio.NewReader(conn)
				data, _, err := reader.ReadLine()
				nickname := string(data)
				if err != nil {
					return
				}

				log.Printf("New client connected %s", nickname)
				allClients[conn] = nickname
				clientCount += 1

				for {
					data, _, err := reader.ReadLine()

					if err != nil {
						break
					}
					messages <- fmt.Sprintf("#%s: %s", nickname, data)

				}

				deadConnections <- conn
			}(conn)

		case message := <-messages:

			target := findTarget(message)

			recipient := 0
			for conn, _ := range allClients {
				go func(conn net.Conn, message string) {

					if isTarget(target, allClients[conn]) {

						_, err := conn.Write([]byte(message + "\n"))
						recipient += 1
						if err != nil {
							deadConnections <- conn
						}
					}

				}(conn, message)
			}
			log.Printf("New message %s", message)
		case conn := <-deadConnections:
			log.Printf("Client %s disconnected", allClients[conn])
			delete(allClients, conn)
		}

	}
}

func isTarget(target string, nickname string) bool {
	if target == "" {
		return true
	}
	if target == nickname {
		return true
	}
	return false
}

func findTarget(msg string) string {
	words := strings.Split(msg, " ")
	for _, word := range words {
		if strings.HasPrefix(word, "@") {
			return strings.TrimPrefix(word, "@")
		}
	}
	return ""
}