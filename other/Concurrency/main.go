package main

import (
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "🐑"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "🐂"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	for {

		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}

func count(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
