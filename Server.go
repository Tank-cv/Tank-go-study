package main

import (
	"bufio"   // 用于读取用户输入和网络数据
	"fmt"     // 打印日志
	"net"     // TCP 网络库
	"strings" // 字符串处理
	"sync"    // 保护共享数据
)

// 全局存储用户信息
var clients = make(map[string]net.Conn) // key:用户名, value:连接对象
var mu sync.Mutex                       // 保护 clients map

func main() {
	// 监听 TCP 端口 8888
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server started on port 12345")

	for {
		// 等待用户连接
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 为每个用户开启一个 goroutine 处理
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// 第一步：要求用户输入用户名
	conn.Write([]byte("Enter your username:\n "))
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username) // 去掉换行符

	// 注册用户
	mu.Lock()
	clients[username] = conn
	mu.Unlock()
	fmt.Println(username, "joined the chat")

	// 提示用户如何发送消息
	conn.Write([]byte("To send message, type: recipient_username:message\n"))

	// 循环读取消息
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(username, "disconnected")
			break
		}
		msg = strings.TrimSpace(msg)
		// 格式必须是 target:message
		parts := strings.SplitN(msg, ":", 2)
		if len(parts) != 2 {
			conn.Write([]byte("Invalid format. Use recipient:message\n"))
			continue
		}
		target, message := parts[0], parts[1]

		// 发送消息给目标用户
		mu.Lock()
		targetConn, ok := clients[target]
		mu.Unlock()
		if ok {
			targetConn.Write([]byte(username + ": " + message + "\n"))
		} else {
			conn.Write([]byte("User " + target + " not found\n"))
		}
	}

	// 用户断开连接后，删除记录
	mu.Lock()
	delete(clients, username)
	mu.Unlock()
}
