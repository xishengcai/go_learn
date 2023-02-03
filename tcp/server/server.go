package main

import (
	"bufio"
	"io/ioutil"
	"net"
	"os/exec"
	"strings"

	"k8s.io/klog/v2"
)

func process(conn net.Conn) {
	// 处理完关闭连接
	defer conn.Close()

	// 针对当前连接做发送和接受操作

	for {
		reader := bufio.NewReader(conn)
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			conn.Write([]byte(err.Error()))
			break
		}
		lineMsg := string(buf[:n])
		opType := strings.Split(lineMsg, ":")

		switch opType[0] {
		case "ftp":
			// ftp: path
			err = downloadFile(opType[1], conn)
			if err != nil {
				break
			}
		default:
			err = runCmd(string(lineMsg), conn)
			if err != nil {
				break
			}
		}
	}
}

func main() {
	// 建立 tcp 服务
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		klog.Infof("listen failed, err:%v", err)
		return
	}

	for {
		// 等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			klog.Infof("accept failed, err:%v", err)
			continue
		}
		// 启动一个单独的 goroutine 去处理连接
		go process(conn)
	}
}

func downloadFile(path string, conn net.Conn) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		klog.Error("read file failed, ", err.Error())
		return err
	}
	n, err := conn.Write(b)
	if err != nil {
		return err
	}
	klog.Infof("download file: %s, size: %d M", path, n/1024/1024)
	return nil
}

func runCmd(command string, conn net.Conn) error {
	var cmd *exec.Cmd
	klog.Infof("command: %s", command)
	cmdSlice := strings.Split(command, " ")
	switch len(command) {
	case 0:
	case 1:
		cmd = exec.Command(cmdSlice[0])
	default:
		cmd = exec.Command(cmdSlice[0], cmdSlice[1:]...)
	}

	var output []byte
	output, err := cmd.CombinedOutput()
	if err != nil {
		_, err = conn.Write([]byte(err.Error()))
		if err != nil {
			klog.Error(err)
			return err
		}
	}

	// 将接受到的数据返回给客户端
	_, err = conn.Write(output)
	if err != nil {
		klog.Infof("write from conn failed, err:%v", err)
	}
	return err
}
