package client

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func ListDir() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", TcpService)
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
	tcpAddr, err := net.ResolveTCPAddr("tcp4", TcpService)
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

//支持断点续传
func DownloadV2(path string) {
	l, err := net.Listen("tcp", TcpService)
	if err != nil {
		log.Fatalf("error listen: %s\n", err)
	}
	defer l.Close()

	log.Println("waiting accept.")
	conn, err := l.Accept()
	if err != nil {
		log.Fatalf("accept faild: %s\n", err)
	}
	fileName := filepath.Base(path)
	DownloadConn(conn, fileName)

}
func writeFile(content []byte, fileName string) {
	if len(content) != 0 {
		fp, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		defer fp.Close()
		if err != nil {
			log.Fatalf("open file faild: %s\n", err)
		}
		_, err = fp.Write(content)
		if err != nil {
			log.Fatalf("append content to file faild: %s\n", err)
		}
		log.Printf("append content: 【%s】 success\n", string(content))
	}
}

func getFileStat(fileName string) int64 {
	fileinfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file size: %d\n", 0)
			return int64(0)
		}
		log.Fatalf("get file stat faild: %s\n", err)
	}
	log.Printf("file size: %d\n", fileinfo.Size())
	return fileinfo.Size()
}

func DownloadConn(conn net.Conn, fileName string) {
	defer conn.Close()
	for {
		var buf = make([]byte, 10)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("server io EOF\n")
				return
			}
			log.Fatalf("server read faild: %s\n", err)
		}
		log.Printf("recevice %d bytes, content is 【%s】\n", n, string(buf[:n]))
		switch string(buf[:n]) {
		case "start-->":
			off := getFileStat(fileName)
			// int conver string
			stringoff := strconv.FormatInt(off, 10)
			info := string(stringoff) + " " + fileName
			_, err = conn.Write([]byte(info))
			if err != nil {
				log.Fatalf("server write faild: %s\n", err)
			}
			continue
		case "<--end":
			log.Fatalf("receive over\n")
			return
		}
		writeFile(buf[:n], fileName)
	}
}
