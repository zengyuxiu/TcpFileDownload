package server

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const service = ":1200"

const addrUdpServer = ":9981"
const addrUdpClient = ":9982"

func checkError(err error) {
	if err != nil {
		log.Errorf("Fatal error: %v", err)
		os.Exit(1)
	}
}
