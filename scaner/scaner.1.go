package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var dict = []string{

	"ROOT.7z",
	"ROOT.bz2",
	"ROOT.gz",
	"ROOT.rar",
	"ROOT.tar.gz",
	"ROOT.tar.bz2",
	"ROOT.zip",
	"ROOT.z",
	"ROOT.tar.z",
	"ROOT.war",
	"wwwroot.7z",
	"wwwroot.bz2",
	"wwwroot.gz",
	"wwwroot.rar",
	"wwwroot.tar.gz",
	"wwwroot.tar.bz2",
	"wwwroot.zip",
	"wwwroot.z",
	"wwwroot.tar.z",
	"wwwroot.war",
	"htdocs.7z",
	"htdocs.bz2",
	"htdocs.gz",
	"htdocs.rar",
	"htdocs.tar.gz",
	"htdocs.tar.bz2",
	"htdocs.zip",
	"htdocs.z",
	"htdocs.tar.z",
	"htdocs.war",
	"www.7z",
	"www.bz2",
	"www.gz",
	"www.rar",
	"www.tar.gz",
	"www.tar.bz2",
	"www.zip",
	"www.z",
	"www.tar.z",
	"www.war",
	"html.7z",
	"html.bz2",
	"html.gz",
	"html.rar",
	"html.tar.gz",
	"html.tar.bz2",
	"html.zip",
	"html.z",
	"html.tar.z",
	"html.war",
	"web.7z",
	"web.bz2",
	"web.gz",
	"web.rar",
	"web.tar.gz",
	"web.tar.bz2",
	"web.zip",
	"web.z",
	"web.tar.z",
	"web.war",
	"webapps.7z",
	"webapps.bz2",
	"webapps.gz",
	"webapps.rar",
	"webapps.tar.gz",
	"webapps.tar.bz2",
	"webapps.zip",
	"webapps.z",
	"webapps.tar.z",
	"webapps.war",
	"public.7z",
	"public.bz2",
	"public.gz",
	"public.rar",
	"public.tar.gz",
	"public.tar.bz2",
	"public.zip",
	"public.z",
	"public.tar.z",
	"public.war",
	"public_html.7z",
	"public_html.bz2",
	"public_html.gz",
	"public_html.rar",
	"public_html.tar.gz",
	"public_html.tar.bz2",
	"public_html.zip",
	"public_html.z",
	"public_html.tar.z",
	"public_html.war",
	"uploads.7z",
	"uploads.bz2",
	"uploads.gz",
	"uploads.rar",
	"uploads.tar.gz",
	"uploads.tar.bz2",
	"uploads.zip",
	"uploads.z",
	"uploads.tar.z",
	"uploads.war",
	"website.7z",
	"website.bz2",
	"website.gz",
	"website.rar",
	"website.tar.gz",
	"website.tar.bz2",
	"website.zip",
	"website.z",
	"website.tar.z",
	"website.war",
	"api.7z",
	"api.bz2",
	"api.gz",
	"api.rar",
	"api.tar.gz",
	"api.tar.bz2",
	"api.zip",
	"api.z",
	"api.tar.z",
	"api.war",
	"test.7z",
	"test.bz2",
	"test.gz",
	"test.rar",
	"test.tar.gz",
	"test.tar.bz2",
	"test.zip",
	"test.z",
	"test.tar.z",
	"test.war",
	"app.7z",
	"app.bz2",
	"app.gz",
	"app.rar",
	"app.tar.gz",
	"app.tar.bz2",
	"app.zip",
	"app.z",
	"app.tar.z",
	"app.war",
	"backup.7z",
	"backup.bz2",
	"backup.gz",
	"backup.rar",
	"backup.tar.gz",
	"backup.tar.bz2",
	"backup.zip",
	"backup.z",
	"backup.tar.z",
	"backup.war",
	"bin.7z",
	"bin.bz2",
	"bin.gz",
	"bin.rar",
	"bin.tar.gz",
	"bin.tar.bz2",
	"bin.zip",
	"bin.z",
	"bin.tar.z",
	"bin.war",
	"bak.7z",
	"bak.bz2",
	"bak.gz",
	"bak.rar",
	"bak.tar.gz",
	"bak.tar.bz2",
	"bak.zip",
	"bak.z",
	"bak.tar.z",
	"bak.war",
	"old.7z",
	"old.bz2",
	"old.gz",
	"old.rar",
	"old.tar.gz",
	"old.tar.bz2",
	"old.zip",
	"old.z",
	"old.tar.z",
	"old.war",
	"Release.7z",
	"Release.bz2",
	"Release.gz",
	"Release.rar",
	"Release.tar.gz",
	"Release.tar.bz2",
	"Release.zip",
	"Release.z",
	"Release.tar.z",
	"Release.war",
	"inetpub.7z",
	"inetpub.bz2",
	"inetpub.gz",
	"inetpub.rar",
	"inetpub.tar.gz",
	"inetpub.tar.bz2",
	"inetpub.zip",
	"inetpub.z",
	"inetpub.tar.z",
	"inetpub.war",
	"index/xyywbf.htm",
}

func main() {
	fmt.Println("====== Scaner ======")
	fmt.Print("请输入目标URL: ")
	var url string

	reader := bufio.NewReader(os.Stdin)
	url, _ = reader.ReadString('\n')
	fmt.Printf("您输入的URL是: %s", url)

	//工作池参数
	var workmax int
	fmt.Println("请输入最大并发数: ")
	fmt.Scanln(&workmax)

	job := make(chan string, len(dict))
	var waitgroup sync.WaitGroup

	//启动工作池
	for i := 0; i < workmax; i++ {
		waitgroup.Add(1)
		go worker(url, job, &waitgroup)

	}

	for _, filename := range dict {
		job <- filename

	}
	close(job)

	waitgroup.Wait()
	fmt.Println("扫描完成！")

}

// 建立一个客户端
var client = &http.Client{
	Timeout: 5 * time.Second,
}

// 工作池函数
func worker(url string, job chan string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	for filename := range job {
		scan(url, filename)
	}
}

// 扫描函数
func scan(url, filename string) {
	//构建完整的URL
	url = strings.TrimSpace(url)
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		url = strings.TrimRight(url, "/")

	} else {
		url = "http://" + strings.TrimRight(url, "/")
	}

	//构建完整的URL
	fullurl := url + "/" + filename

	//发送HTTP请求
	resp, err := http.NewRequest("GET", fullurl, nil)
	if err != nil {
		fmt.Printf("构造请求 %s 失败: %v\n", fullurl, err)
		return
	}
	resp.Header.Set("User-Agent", "Mozilla/5.0")

	response, err := client.Do(resp)
	if err != nil {
		fmt.Printf("请求 %s 失败: %v\n", fullurl, err)
		return
	}
	defer response.Body.Close()

	//检查响应状态码
	if response.StatusCode == 200 || response.StatusCode == 403 {
		if response.ContentLength > 1024 {
			fmt.Printf("发现文件: %s (状态码: %d)\n", fullurl, response.StatusCode)
		}
	}

	//读取文件内容
	/*if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("读取响应体失败: %v\n", err)
			return
		}

		if len(body) > 0 {
			fmt.Printf("文件内容: %s\n", string(body))
		}
	}*/
	//容易爆内存不建议读取大文件内容，可以根据需要调整读取逻辑，例如只读取前几行或前几KB的内容。

}
