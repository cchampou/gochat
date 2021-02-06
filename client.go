package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	c, err := net.Dial("tcp", ":6000")

	messages := make(chan string)

	inputs := make(chan string)

	if err != nil {
		log.Fatal(err)
	}

	go func(c net.Conn) {
		reader := bufio.NewReader(c)
		for {
			data, _ := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				break
			}
			messages <- data
		}
	}(c)

	go func(c net.Conn) {
		reader := bufio.NewReader(os.Stdin)
		for {
			data, _ := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				break
			}
			inputs <- data
		}
	}(c)

	for {
		select {
		case message := <-messages:
			fmt.Print(message)
		case input := <-inputs:
			go func(c net.Conn, input string) {
				_, err := c.Write([]byte(input))
				if err != nil {
					log.Fatal(err)
				}
			}(c, input)
		}
	}
}
