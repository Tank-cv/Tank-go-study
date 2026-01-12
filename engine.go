package main

/*import "fmt"
func main() {
    for i := 1; i <= 1024; i++ {
        address := fmt.Sprintf("scanme.nmap.org:%d", i)
        fmt.Println(address)
    }
}*/
/*import (
    "fmt"
    "net"
)
func main() {
    for i := 1; i <= 1024; i++ {
        address := fmt.Sprintf("scanme.nmap.org:%d", i)
        conn, err := net.Dial("tcp", address)
        if err != nil {
            // port is closed or filtered.
continue }
        conn.Close()
        fmt.Printf("%d open\n", i)
} }*/
/*import (
    "fmt"
    "net"
)
func main() {
    for i := 1; i <= 1024; i++ {
        go func(j int) {
            address := fmt.Sprintf("scanme.nmap.org:%d", j)
            conn, err := net.Dial("tcp", address)
            if err != nil {
                return
            }
            conn.Close()
            fmt.Printf("%d open\n", j)
        }(i)
    }
}*/

/*import (
    "fmt"
    "net"
"sync"
)
func main() {
  var wg sync.WaitGroup
    for i := 1; i <= 1024; i++ {
      wg.Add(1)
        go func(j int) {
          defer wg.Done()
            address := fmt.Sprintf("127.0.0.1:%d", j)
            conn, err := net.Dial("tcp", address)
            if err != nil {
                return
            }
            conn.Close()
            fmt.Printf("%d open\n", j)
        }(i)
    }
  wg.Wait()
}*/

/*import (
    "fmt"
    "sync"
)
func worker(ports chan int, wg *sync.WaitGroup) {
     for p := range ports {
        fmt.Println(p)
        wg.Done()
    }
}
func main() {
    ports := make(chan int, 100)
    var wg sync.WaitGroup
    for i := 0; i < cap(ports); i++ {
        go worker(ports, &wg)
    }
    for i := 1; i <= 1024; i++ {
        wg.Add(1)
        ports <- i
    }
    wg.Wait()
    close(ports)
}*/

import (
	"fmt"
	"net"
	"sync"
)

func worker(ports chan int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(ports, results, &wg)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
		close(ports)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		fmt.Printf("%d open\n", port)
	}
}
