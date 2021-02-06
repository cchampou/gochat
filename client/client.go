package main

import (
	"bufio"
	"cchampou.me/network"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	messages := make(chan string)

	inputs := make(chan string)

	c := network.DialServer(6000)

	go func(c net.Conn) {
		reader := bufio.NewReader(c)
		for {
			data, err := reader.ReadString('\n')
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
			data, err := reader.ReadString('\n')
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
