package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var clients = make(map[string]net.Conn)

var shuo sync.Mutex

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("请输入你的名字：\n"))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取名字失败", err)
		return
	}
	name := strings.TrimSpace(string(buf[:n]))
	shuo.Lock()
	clients[name] = conn
	shuo.Unlock()

	fmt.Println(name, "加入了聊天")

	conn.Write([]byte("发送格式：名字:消息内容\n"))
	conn.Write([]byte("输入exit退出聊天\n"))

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取消息失败", err)
			return
		}

		xiaoxi := strings.TrimSpace(string(buf[:n]))

		if xiaoxi == "exit" {
			fmt.Println(name, "退出了聊天")
			shuo.Lock()
			delete(clients, name)
			shuo.Unlock()
			return
		}

		part := strings.SplitN(xiaoxi, ":", 2)
		if len(part) != 2 {
			conn.Write([]byte("消息格式错误，请使用名字:消息内容的格式\n"))
			continue
		}
		mubuao := part[0]
		neirong := part[1]

		shuo.Lock()
		targetConn, ok := clients[mubuao]
		shuo.Unlock()

		if ok {
			targetConn.Write([]byte(name + "对你说：" + neirong + "\n"))
		}
	}
}
func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":12345")

	if err != nil {
		fmt.Println("启动失败", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器已启动，等待连接...")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败", err)
			return
		}

		fmt.Println("新连接来自", conn.RemoteAddr())
		fmt.Fprintf(conn, "欢迎连接到服务器！\n")

		go handleConnection(conn)
	}

}
