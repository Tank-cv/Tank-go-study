package main

import (
	"fmt"
	"net"
)

// 垃圾代码生产商  垃圾+1  ’（-_-）‘
func main() {
	fmt.Println("||========极简抓包========||")
	var host string
	fmt.Println("请输入网址")
	fmt.Scanln(&host)

	fmt.Println("请输入端口")
	var port string
	fmt.Scanln(&port)

	target := host + ":" + port

	conn, err := net.Dial("tcp", target)
	if err != nil {
		fmt.Printf("连接目标失败：%v\n", err)
		return
	}
	defer conn.Close()

	req := "GET / HTTP/1.1\r\nHost: " + host + "\r\n"
	_, err = conn.Write([]byte(req))
	if err != nil {
		fmt.Printf("发送请求失败：%v\n", err)
		return
	}

	var Data []byte
	buf := make([]byte, 100000)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			Data = append(Data, buf[:n]...)
		}
		if err != nil {
			fmt.Printf("读取数据失败：%v\n", err)
			break
		}
	}

	k := 0
	var counnt int = 0
	for range Data {

		if k+3 >= len(Data) {
			break
		}

		if Data[k] == 'f' && Data[k+1] == 'l' && Data[k+2] == 'a' && Data[k+3] == 'g' {
			fmt.Println("\n== 找到flag ==")

			if k+4 < len(Data) && Data[k+4] == '{' {
				j := k + 5

				for j < len(Data) && Data[j] != '}' {
					j++
				}

				if j < len(Data) {
					fmt.Println("===========flag===========")
					fmt.Println("flag{", string(Data[k+5:j]), "}")
					fmt.Println("=========================")
					counnt = 1
					break
				}
			}
		}
		k += 1
	}

	if counnt == 0 {
		fmt.Println("\n== 未找到flag ==")
	}

	fmt.Println("\n== 从目标读取的数据包内容 ==")
	fmt.Println(string(Data))
}
