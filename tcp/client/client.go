package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"k8s.io/klog/v2"
)

func main() {
	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		klog.Errorf("conn server failed, err:%v", err)
		return
	}
	go func() {
		for {
			// 从服务端接收回复消息
			var buf [1024]byte
			n, err := conn.Read(buf[:])
			if err != nil {
				klog.Errorf("read failed:%v", err)
				return
			}
			fmt.Print(string(buf[:n]))
		}

	}()

	// 2、使用 conn 连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			klog.Errorf("send failed, err:%v", err)
			return
		}
	}
}

func saveFile() {

}
