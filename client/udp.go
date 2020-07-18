package client

import (
	"fmt"
	"net"
)

func UdpClient() {
	addr_server, err := net.ResolveUDPAddr("udp4", addrUdpServer)
	checkError(err)
	conn, err := net.DialUDP("udp4", nil, addr_server)
	checkError(err)
	for {
		var message string
		fmt.Print("< ")
		_, err = fmt.Scanf("%s\n", &message)
		_, err = conn.Write([]byte(message))
		checkError(err)
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Print("> ")
		fmt.Println(string(buf[0:n]))
	}

}
