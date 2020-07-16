package server

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
)

const service = ":1200"

func RunServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	request := make([]byte, 128)
	defer conn.Close()

	for {
		length, err := conn.Read(request)
		if err != nil {
			log.Errorf("Read Error %v", err)
			break
		}
		if length == 0 {
			break
		} else if strings.TrimSpace(string(request[:length])) == "tree" {
			out, err := exec.Command("bash", "-c", "tree").Output()
			if err != nil {
				conn.Write([]byte(err.Error()))
				break
			} else {
				conn.Write([]byte(out))
				break
			}
		} else {
			path := strings.TrimSpace(string(request[:length]))
			context, err := ioutil.ReadFile(path)
			if err != nil {
				//conn.Write([]byte(err.Error()))
				log.Errorf("Fatal error: %v", err)
				break
			} else {
				conn.Write(context)
				break
			}
		}
	}
	request = make([]byte, 128)
}
func checkError(err error) {
	if err != nil {
		log.Errorf("Fatal error: %v", err)
		os.Exit(1)
	}
}
