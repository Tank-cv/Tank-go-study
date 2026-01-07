package main

import (
	"fmt"
	"sync"
)

var cashier1chan = make(chan int, 10000)
var cashier2chan = make(chan int, 10000)
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
	for id := range cashier1chan {
		var productName string
		fmt.Printf("顾客%d 到达收银台1，请输入商品名称: ", id)
		fmt.Scanln(&productName)
		price, ok := productMap[productName]
		if !ok {
			fmt.Printf("收银台1：未找到商品 %s，价格视为0\n", productName)
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
	for id := range cashier2chan {
		var productName string
		fmt.Printf("顾客%d 到达收银台2，请输入商品名称: ", id)
		fmt.Scanln(&productName)
		price, ok := productMap[productName]
		if !ok {
			fmt.Printf("收银台2：未找到商品 %s，价格视为0\n", productName)
			continue
		}
		total += price
		fmt.Printf("收银台2：商品 %s 价格 ¥%d\n", productName, price)
	}
	fmt.Printf("收银台2总计: %d\n", total)
	wait.Done()
}
func produce(ch chan<- int) {
	// 模拟一定数量的顾客到达，只将顾客发送到目标收银台的 channel
	for i := 1; i <= 10000; i++ {
		fmt.Printf("客人%d 到达队列\n", i)
		ch <- i
	}
	close(ch)
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
		go produce(cashier1chan)
	case 2:
		wait.Add(1)
		go cashier2()
		go produce(cashier2chan)
	default:
		fmt.Println("无效选择，退出")
		return
	}
	wait.Wait()

}
