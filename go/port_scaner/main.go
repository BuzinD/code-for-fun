package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	successPorts := make([]int, 0, 100)
	for port := 0; port <= 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
			if err != nil {
				return
			}
			conn.Close()
			successPorts = append(successPorts, port)
		}(port)

	}

	wg.Wait()

	fmt.Println(successPorts)

}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}
