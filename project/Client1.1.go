package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println("连接失败", err)
	}

	defer conn.Close()

	fmt.Println("连接成功")

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("读取消息失败", err)
				return
			}
			fmt.Print(string(buf[:n]))
		}
	}()
	/*for {
		var input string
		fmt.Scanln(&input)
		conn.Write([]byte(input))

	}*/
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		conn.Write([]byte(input))
	}
}
