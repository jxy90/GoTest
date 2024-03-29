package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func main() {
	AsyncPool0()
}

func worker(ports chan int, result chan int) {
	for port := range ports {
		address := fmt.Sprintf("10,202,101,31:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//fmt.Printf("%v Closed! \n", address)
			result <- 0
			//wg.Done()
			continue
		}
		conn.Close()
		//fmt.Printf("%v Open! \n", address)
		result <- port
		//wg.Done()

	}
}

func AsyncPool() {
	//var wg sync.WaitGroup
	ports := make(chan int, 100)
	results := make(chan int, 10)
	var openPorts, closePorts []int
	for i := 0; i < 100; i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 21; i < 100; i++ {
			ports <- i
		}
	}()
	for i := 21; i < 100; i++ {
		port := <-results
		if port == 0 {
			closePorts = append(closePorts, port)
		} else {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	sort.Ints(closePorts)
	//wg.Wait()
	fmt.Println(openPorts)
	fmt.Println(closePorts)
}

func AsyncPool0() {
	ports := make(chan int, 100)
	results := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go worker0(ports, results, &wg)
	}
	go func() {
		for i := 0; i < 150; i++ {
			//wg.Add(1)
			ports <- i
		}
	}()
	nums := make([]int, 0)
	for i := 0; i < 150; i++ {
		result := <-results
		if result != 0 {
			nums = append(nums, result)
		}
	}
	//wg.Wait()
	fmt.Println(nums)
}

func worker0(ports, result chan int, wg *sync.WaitGroup) {
	for port := range ports {
		address := fmt.Sprintf("192.168.0.1:%v", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			//fmt.Printf("%v Error! \n", address)
			//wg.Done()
			continue
		}
		result <- port
		conn.Close()
		//fmt.Printf("%v Open! \n", address)
		//wg.Done()
	}
}

func Async() {
	var wg sync.WaitGroup
	for i := 21; i < 120; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", index)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%v Closed! \n", address)
				return
			}
			conn.Close()
			fmt.Printf("%v Open! \n", address)
		}(i)

	}
	wg.Wait()
}

func Sync() {
	for i := 21; i < 120; i++ {
		address := fmt.Sprintf("20.194.168.28:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%v Closed!", address)
			continue
		}
		conn.Close()
		fmt.Printf("%v Open!", address)
	}
}
