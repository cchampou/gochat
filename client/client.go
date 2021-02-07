package main

import (
	"cchampou.me/customIo"
	"fmt"
	"net"
	"os"
)

func pushMessages(c net.Conn, messages chan string) {
	reader := customIo.CreateReader(c)
	for {
		data := customIo.ReadLine(reader)
		messages <- data
	}
}

func pushStdin(inputs chan string) {
	reader := customIo.CreateReader(os.Stdin)
	for {
		data := customIo.ReadLine(reader)
		inputs <- data
	}
}

func main() {

	customIo.ClearWindow()

	messages := make(chan string)

	inputs := make(chan string)

	c := customIo.DialServer(6000)

	go pushMessages(c, messages)

	go pushStdin(inputs)

	for {
		select {
		case message := <-messages:
			fmt.Print(message)
		case input := <-inputs:
			go func(c net.Conn, input string) {
				customIo.WriteString(c, input)
			}(c, input)
		}
	}
}
