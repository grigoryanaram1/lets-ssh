package main

import (
	"net"
	"os"
	"strconv"
	"strings"
)



func SSHEstablishTcpConnection(ip string, port int) net.Conn {
	servAddr := ip + ":" + strconv.Itoa(port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	return conn
}

func ValidateVersionExchangeMessage(message string) {
	chunks := strings.Split(message, "-")
	if len(chunks) < 3 {
		println("Invalid version message received.")
		os.Exit(1)
	}
	// Rule 1 : check first 3 bytes, must be "SSH"
	if strings.Compare(chunks[0], "SSH") != 0 {
		println("Invalid version message received.")
		os.Exit(1)
	}
	for
}

func SSHProtocolVersionExchange(conn net.Conn) {
	reply := make([]byte, 255)
	_, err := conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	println("Protocol Version Exchange message : ", string(reply))
}

func SSHCloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		println("Close connection failed:", err.Error())
	}
}

func connectToServer(ip string, port int) {
	conn := SSHEstablishTcpConnection(ip, port)
	SSHProtocolVersionExchange(conn)
	SSHCloseConnection(conn)
}

func main() {
	connectToServer("192.168.64.3", 22)
}
