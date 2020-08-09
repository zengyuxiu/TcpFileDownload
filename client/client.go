package client

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const TcpService = ":1200"
const addrUdpServer = ":9981"
const addrUdpClient = ":9982"

func HandleTcp(list bool, path string) {
	if list == false {
		if path != "" {
			log.Infof("Path : %v", path)
			DownloadV2(path)
		}
	} else {
		ListDir()
	}
}
func HandleUdp() {
	UdpClient()
}
func checkError(err error) {
	if err != nil {
		log.Errorf("Fatal error: %v", err)
		os.Exit(1)
	}
}
