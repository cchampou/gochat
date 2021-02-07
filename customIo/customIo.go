package customIo

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
)

func ToServerAddr(port int) string {
	address := ":" + strconv.Itoa(port)
	return address
}

func DialServer(port int) net.Conn {
	c, err := net.Dial("tcp", ToServerAddr(port))
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func CreateServer(port int) net.Listener {
	server, err := net.Listen("tcp", ToServerAddr(port))

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

func ReadLine(reader *bufio.Reader) string {
	data, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func CreateReader(c io.Reader) *bufio.Reader {
	reader := bufio.NewReader(c)
	return reader
}

func ClearWindow() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
