package main

import (
	"fmt"
	"sync"
)

var cashier1chan = make(chan string, 10000)
var cashier2chan = make(chan string, 10000)
var wait sync.WaitGroup

var productMap = map[string]int{
	"商品1":  10,
	"商品2":  20,
	"商品3":  15,
	"商品4":  25,
	"商品5":  30,
	"商品6":  18,
	"商品7":  22,
	"商品8":  28,
	"商品9":  12,
	"商品10": 35,
}

type producemap struct {
	name  string
	price int
}

func cashier1() {
	var total int
	println("请输入商品名称，输入exit退出")
	for{
		var input string
		fmt.Scanln(&input)
		cashier1chan<-input
		if input=="exit"{
			break
		}
	}
	for id := range cashier1chan {
		var productName string
		fmt.Printf("收银台1，商品名称: %s", id)
		fmt.Scanln(&productName)
		price, ok := productMap[productName]
		if !ok {
			continue
		}
		total += price
		fmt.Printf("收银台1：商品 %s 价格 ¥%d\n", productName, price)
	}
	fmt.Printf("收银台1总计: %d\n", total)
	wait.Done()
}
func cashier2() {
	var total int
	println("请输入商品名称，输入exit退出")
	for{
		var input string
		fmt.Scanln(&input)
		cashier1chan<-input
		if input=="exit"{
			break
		}
	}
	for id := range cashier1chan {
		var productName string
		fmt.Printf("收银台2，商品名称: %s", id)
		fmt.Scanln(&productName)
		price, ok := productMap[productName]
		if !ok {
			continue
		}
		total += price
		fmt.Printf("收银台2：商品 %s 价格 ¥%d\n", productName, price)
	}
	fmt.Printf("收银台2总计: %d\n", total)
	wait.Done()
}
func main() {
	println("====收银系统====")
	println("1.收银台1 2.收银台2")

	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		wait.Add(1)
		go cashier1()
		
	case 2:
		wait.Add(1)
		go cashier2()
		
	default:
		fmt.Println("无效选择，退出")
		return
	}
	wait.Wait()

}
