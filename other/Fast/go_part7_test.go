package Fast

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
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
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	fmt.Println("put ", id)
	c <- id
}

func Test30_2(t *testing.T) {
	//地鼠工厂gopher
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go gopherId(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("get ", gopherID)
		case <-timeout:
			fmt.Println("time out!!!")
			return
		}

	}
	fmt.Println("----------")
}

func Test30_3(t *testing.T) {
	//地鼠工厂gopher
	c1, c2 := make(chan string), make(chan string)
	go sourceGopher(c1)
	go filterGopher(c1, c2)
	printGopher(c2)
	fmt.Println("----------")
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}

func sourceGopher(downstream chan string) {
	strs := []string{"hello word", "bad apple", "good bye"}
	for _, str := range strs {
		downstream <- str
	}
	//downstream <- ""
	close(downstream)
}

func filterGopher(upStream, downstream chan string) {
	for {
		v, ok := <-upStream
		if !ok {
			close(downstream)
			return
		}
		if !strings.Contains(v, "bad") {
			downstream <- v
		}
	}
}

//func filterGopher(upStream, downstream chan string) {
//	for item := range upStream {
//		if !strings.Contains(item, "bad") {
//			downstream <- item
//		}
//	}
//	close(downstream)
//}

//func printGopher(upstream chan string) {
//	for {
//		v, ok := <-upstream
//		if !ok {
//			return
//		}
//		fmt.Println(v)
//	}
//}
func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}

func Test31(t *testing.T) {
	//地鼠工厂gopher
	visited := Visited{
		mu:      sync.Mutex{},
		visited: map[string]int{},
	}
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	go visited.VisitedLink("123")
	time.Sleep(time.Second * 5)
	fmt.Println(visited.visited)
	fmt.Println("----------")
}

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitedLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
