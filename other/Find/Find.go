package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var findStr = "test"
var matches = 0

var workerCount = 0
var maxWorkerCount = 8
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var findMatch = make(chan bool)

func main() {
	//c := make(chan int)
	start := time.Now()
	workerCount = 1
	go search("/Users/jiangxiaoyu/", true)
	waitForWorkers()
	fmt.Println("matches", matches)
	fmt.Println(time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-findMatch:
			matches++
		}
	}
}

//matches 18298
//4.480994123s
//matches 18298
//24.183085826s
func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if strings.Contains(name, findStr) {
				findMatch <- true
			}
			if file.IsDir() {
				newPath := path + name + "/"
				if workerCount < maxWorkerCount {
					searchRequest <- newPath
				} else {
					search(newPath, false)
				}
			}
		}
	}
	if master {
		workerDone <- true
	}
}
