package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Trying to connect to server 127.0.0.1:12345...")
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server!")

	reader := bufio.NewReader(conn)
	stdin := bufio.NewReader(os.Stdin)

	// 读取服务器提示输入用户名
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Server disconnected:", err)
		return
	}
	fmt.Print(msg) // 显示服务器提示

	// 输入用户名
	username, _ := stdin.ReadString('\n')
	username = strings.TrimSpace(username)
	conn.Write([]byte(username + "\n"))

	fmt.Println("You can now send messages in format: recipient_username:message")

	// 启动 goroutine 接收服务器消息
	go func() {
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nDisconnected from server")
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	// 循环读取用户输入发送消息
	for {
		fmt.Print("> ") // 提示符
		input, _ := stdin.ReadString('\n')
		input = strings.TrimSpace(input)
		if input != "" {
			conn.Write([]byte(input + "\n"))
		}
	}
}
