package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	id     int
	title  string
	author string
	count  int
}

type Library struct {
	Books map[int]*Book
}

//func NewLibrary() *Library {
//	return &Library{
//		Books: make(map[int]*Book),
//	}
//}

func (l *Library) Addbook(id int, title, author string, count int) error {
	if l.Books == nil {
		l.Books = make(map[int]*Book) // 添加图书
	}
	l.Books[id] = &Book{
		id:     id,
		title:  title,
		author: author,
		count:  count,
	}
	return nil
}

func (l *Library) Searchbook(id int) (*Book, bool) { // 搜索图书
	if book, exists := l.Books[id]; exists {
		return book, true
	}
	return nil, false
}

func (l *Library) Listbook() { // 列出所有图书
	if len(l.Books) == 0 {
		fmt.Println("图书馆为空")
		return
	}
	fmt.Println("=====图书列表======")
	for id, book := range l.Books {
		fmt.Printf("ID:%d, Title:%s, Author:%s, Count:%d\n", id, book.title, book.author, book.count)
	}

}

func (l *Library) Updatebook(id int, title, author string, count int) error { // 更新图书信息
	if book, exists := l.Books[id]; exists {
		book.title = title
		book.author = author
		book.count = count
		return nil
	}
	return fmt.Errorf("book with ID %d not found", id)
}

func (l *Library) Deletebook(id int) error { // 删除图书
	if _, exists := l.Books[id]; exists {
		delete(l.Books, id)
		return nil
	}
	return fmt.Errorf("book with ID %d not found", id)
}
func main() {
	lib := &Library{}
	// 预置两本书用于测试
	lib.Addbook(1, "1984", "George Orwell", 5)
	lib.Addbook(2, "To Kill a Mockingbird", "Harper Lee", 3)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n====进入图书管理系统====")
		fmt.Println("1. 添加  2. 搜索  3. 列出  4. 更新  5. 删除  0. 退出")
		fmt.Print("请选择操作: ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		switch line {
		case "1":
			fmt.Print("输入 ID: ")
			id := readInt(reader)
			fmt.Print("输入 Title: ")
			title := readLine(reader)
			fmt.Print("输入 Author: ")
			author := readLine(reader)
			fmt.Print("输入 Count: ")
			count := readInt(reader)
			lib.Addbook(id, title, author, count)
			fmt.Println("添加成功")
		case "2":
			fmt.Print("输入要搜索的 ID: ")
			id := readInt(reader)
			if book, ok := lib.Searchbook(id); ok {
				fmt.Printf("找到: ID:%d, Title:%s, Author:%s, Count:%d\n", book.id, book.title, book.author, book.count)
			} else {
				fmt.Println("未找到该图书")
			}
		case "3":
			lib.Listbook()
		case "4":
			fmt.Print("输入要更新的 ID: ")
			id := readInt(reader)
			fmt.Print("新的 Title: ")
			title := readLine(reader)
			fmt.Print("新的 Author: ")
			author := readLine(reader)
			fmt.Print("新的 Count: ")
			count := readInt(reader)
			if err := lib.Updatebook(id, title, author, count); err != nil {
				fmt.Println("更新失败:", err)
			} else {
				fmt.Println("更新成功")
			}
		case "5":
			fmt.Print("输入要删除的 ID: ")
			id := readInt(reader)
			if err := lib.Deletebook(id); err != nil {
				fmt.Println("删除失败:", err)
			} else {
				fmt.Println("删除成功")
			}
		case "0":
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("无效选项，请重试")
		}
	}
}

// 辅助读取函数
func readLine(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(r *bufio.Reader) int {
	for {
		s := readLine(r)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Print("输入不是有效整数，请重新输入: ")
			continue
		}
		return n
	}
}

//cd c:\Users\hr\Desktop\go
//go run .\Untitled-1.go
