package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()

	fmt.Println("=== 已连接服务端 ===")

	// 接收协程
	go func() {
		buf := make([]byte, 1024)
		for {
			n, _ := conn.Read(buf)
			fmt.Println("\n收到服务端：", string(buf[:n]))
		}
	}()

	// 发送
	var input string
	for {
		fmt.Print("请输入：")
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("退出聊天")
			return
		}
		conn.Write([]byte(input))
	}
}
