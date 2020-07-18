package server

import (
	"fmt"
	"net"
)

func handleUdpClient(conn *net.UDPConn) {
	var buf [512]byte
	fmt.Print("> ")
	n, addr, err := conn.ReadFromUDP(buf[0:])
	fmt.Println(string(buf[0:n]))
	if err != nil {
		return
	}
	var message string
	fmt.Print("< ")
	_, err = fmt.Scanf("%s\n", &message)
	conn.WriteToUDP([]byte(message), addr)
}
func UdpServer() {
	addr_server, err := net.ResolveUDPAddr("udp4", addrUdpServer)
	checkError(err)
	conn, err := net.ListenUDP("udp4", addr_server)
	checkError(err)
	defer conn.Close()
	for {
		handleUdpClient(conn)
	}

}
