package server

import (
	"net"
	"time"
)

func handleUdpClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
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
