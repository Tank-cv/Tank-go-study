package main

import (
	"fmt"
	"net"
)

func main() {

	target := "node4.anna.nssctf.cn:25939"

	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("连接目标失败：%v\n", err)
		return
	}
	defer conn.Close()

	req := "GET /index.php HTTP/1.1\r\nHost: node4.anna.nssctf.cn\r\nConnection: close\r\n\r\n"
	_, err = conn.Write([]byte(req))
	if err != nil {
		fmt.Printf("发送请求失败：%v\n", err)
		return
	}

	var Data []byte
	buf := make([]byte, 10000)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			Data = append(Data, buf[:n]...)
		}
		if err != nil {
			break
		}
	}

	fmt.Println("\n=== 从目标读取的数据包内容 ===")
	fmt.Println(string(Data))

}
