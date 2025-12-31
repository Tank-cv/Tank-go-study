package main

import (
	"encoding/base64"
	"encoding/hex"
	f "fmt"
	"net/url"
)

// 先将要调配的函数写出来

func menu() int {
	var ch int
	f.Println("1编码，2解码，0退出")
	f.Scan(&ch)
	return ch

}
func encmenu() int {
	var ch int
	f.Println("编码：1,base64 2,url 3,hex 0退出")
	f.Scan(&ch)
	return ch
}
func decmenu() int {
	var ch int
	f.Println("解码：1,base64 2,url 3,hex 0退出")
	f.Scan(&ch)
	return ch
}
func base64enc(str string) string {
	var b []byte = []byte(str)
	return base64.StdEncoding.EncodeToString(b)
}
func base64dec(str string) (string, error) {
	var d, err = base64.StdEncoding.DecodeString(str)
	return string(d), err
}
func urleco(str string) string {
	return url.QueryEscape(str)

}
func urldeco(str string) (string, error) {
	return url.QueryUnescape(str)
}
func hexenc(str string) string {
	var b []byte = []byte(str)
	return hex.EncodeToString(b)
}
func hexdec(str string) (string, error) {
	var d, err = hex.DecodeString(str)
	return string(d), err
}

// 进入主程序，开始写逻辑
func main() {
	for {
		choice := menu()
		switch choice {
		case 1:
			for {
				ech := encmenu()
				if ech == 0 {
					f.Println("返回主菜单")
					break
				}
				var input string
				f.Println("请输入")
				f.Scanln(&input)
				var ans string
				switch ech {
				case 1:

					ans = base64enc(input)
				case 2:
					ans = urleco(input)
				case 3:
					ans = hexenc(input)
				}
				f.Printf("%s\n", ans)

			}
		case 2:
			for {
				dec := decmenu()
				if dec == 0 {
					f.Println("返回主菜单")
					break
				}
				var input string
				f.Println("请输入")
				f.Scanln(&input)
				ans, err := "", error(nil)
				switch dec {
				case 1:

					ans, err = base64dec(input)
				case 2:
					ans, err = urldeco(input)
				case 3:
					ans, err = hexdec(input)
				}
				if err != nil {
					f.Println("解码错误", err)
				} else {
					f.Println("解码成功", ans)
				}
			}
		default:
			f.Println("程序退出")
			return
		}
	}
}
