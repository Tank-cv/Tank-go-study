package main

import (
	"fmt"
	"io"
	"net"
)

// 错误处理
func errs(err error) {
	if err != nil {
		fmt.Println("错误原因err:", err)
		return
	} else if err == io.EOF {
		fmt.Println("客户端已断开连接")
		return
	}
}

// 接收消息
func JIESOUhandle(conn *net.TCPConn) {

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		errs(err)

		fmt.Println("收到消息:", string(buf[:n]))
	}
}

// 发送消息
func FASONGhandle(conn *net.TCPConn) {

	var input string
	for {
		fmt.Print("请输入要发送的消息:")
		fmt.Scanln(&input)

		buf := []byte(input)

		if input == "exit" {
			fmt.Println("退出聊天")
			return
		}
		n, err := conn.Write(buf)
		errs(err)

		fmt.Println("发送消息:", string(buf[:n]))
	}
}

func main() {
	adrr, err := net.ResolveTCPAddr("tcp", ":12345")
	errs(err)

	listener, err := net.ListenTCP("tcp", adrr)
	errs(err)

	defer listener.Close()

	fmt.Println("服务器已启动，等待连接...")

	for {
		conn, err := listener.AcceptTCP()
		errs(err)
		fmt.Println("连接成功:", conn.RemoteAddr())

		go JIESOUhandle(conn)
		go FASONGhandle(conn)
	}
}
