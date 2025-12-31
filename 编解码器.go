package main

import (
	 "encoding/base64"
	 "encoding/hex"
	f "fmt"
	 "net/url"
	 "strings"
)

func menu() int {
	var ch int
	f.Println("1编码，2解码，eixt退出")
	f.Scan(&ch)
	return ch

}
func encmenu() int {
	var ch int
	f.Println("编码：1,base64 2,url 3,hex eixt退出")
	f.Scan(&ch)
	return ch
}
func decmenu() int {
	var ch int
	f.Println("解码：1,base64 2,url 3,hex eixt退出")
	f.Scan(&ch)
	return ch
}
func base64enc(str string) string {
	var []byte
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func base64dec(str string) (string,error) {
	var d,err =base64.StdEncoding.DecodeString(str)
	return string(d), err
}
func urleco(str string) string {
	var str string
	return url.QueryEscape(str)

}
func urldeco(str string)(string,error) {
	var str string
	return url.QueryUnescape(str)
}
func hexenc(str string) string {
	var []byte
	return hex.EncodeToString([]byte(str))
}
func hexdec(str string) (string,error) {
	var	d,err =hex.DecodeString(str)
}