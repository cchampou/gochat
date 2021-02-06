package network

import (
	"log"
	"net"
	"strconv"
)

func CreateServer(port int) net.Listener {
	address := ":" + strconv.Itoa(port)
	log.Print(address)
	server, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	return server
}

func AcceptConn(server net.Listener) net.Conn {
	conn, err := server.Accept()

	if err != nil {
		panic(err)
	}

	return conn
}

func WriteString(c net.Conn, str string) error {
	_, err := c.Write([]byte(str))
	return err
}
