package network

import (
	"log"
	"net"
	"strconv"
)

func toServerAddr(port int) string {
	address := ":" + strconv.Itoa(port)
	return address
}

func DialServer(port int) net.Conn {
	c, err := net.Dial("tcp", toServerAddr(port))
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func CreateServer(port int) net.Listener {
	server, err := net.Listen("tcp", toServerAddr(port))

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
