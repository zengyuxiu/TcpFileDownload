package client

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

const service = ":1200"

func ListDir() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("tree"))
	checkError(err)
	context, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(context))
}
func Download(path string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte(path))
	checkError(err)
	context, err := ioutil.ReadAll(conn)
	checkError(err)
	err = ioutil.WriteFile(filepath.Base(path), context, 0644)
	checkError(err)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		log.Errorf("Fatal error: %v", err)
		os.Exit(1)
	}
}
