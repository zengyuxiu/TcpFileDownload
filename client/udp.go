package client

import (
	"fmt"
	"net"
	"os"
)

func UdpClient() {
	addr_server, err := net.ResolveUDPAddr("udp4", addrUdpServer)
	checkError(err)
	conn, err := net.DialUDP("udp4", nil, addr_server)
	checkError(err)
	_, err = conn.Write([]byte("time"))
	checkError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
