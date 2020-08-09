package server

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func TcpServer() {
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
		args := strings.Split(string(request), " ")
		if len(args) == 2 {
			conn, err := net.DialTimeout("tcp", service, time.Second*10)
			if err != nil {
				log.Fatalf("client dial faild: %s\n", err)
			}
			bkConn(conn, args)
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

func bkWrite(conn net.Conn, data []byte) {
	_, err := conn.Write(data)
	if err != nil {
		log.Fatalf("send 【%s】 content faild: %s\n", string(data), err)
	}
	log.Printf("send 【%s】 content success\n", string(data))
}

func bkConn(conn net.Conn, args []string) {
	defer conn.Close()

	// 发送"start-->"消息通知服务端，我要开始发送文件内容了
	// 你赶紧告诉我你那边已经接收了多少内容,我从你已经接收的内容处开始继续发送
	bkWrite(conn, []byte("start-->"))
	off, _ := strconv.Atoi(args[0])
	file := args[1]

	// send file content
	fp, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatalf("open file faild: %s\n", err)
	}
	defer fp.Close()

	// set file seek
	// 设置从哪里开始读取文件内容
	_, err = fp.Seek(int64(off), 0)
	if err != nil {
		log.Fatalf("set file seek faild: %s\n", err)
	}
	log.Printf("read file at seek: %d\n", off)

	for {
		// 每次发送10个字节大小的内容
		data := make([]byte, 10)
		n, err := fp.Read(data)
		if err != nil {
			if err == io.EOF {
				// 如果已经读取完文件内容
				// 就发送'<--end'消息通知服务端，文件内容发送完了
				time.Sleep(time.Second * 1)
				bkWrite(conn, []byte("<--end"))
				log.Println("send all content, now quit")
				break
			}
			log.Fatalf("read file err: %s\n", err)
		}
		// 发送文件内容到服务端
		bkWrite(conn, data[:n])
	}
}
