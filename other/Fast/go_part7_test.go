package Fast

import (
	"fmt"
	"testing"
	"time"
)

func Test30(t *testing.T) {
	//地鼠工厂gopher
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go gopherId(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("get ", gopherID)
	}
	fmt.Println("----------")
}

func gopherId(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("put ", id)
	c <- id
}
