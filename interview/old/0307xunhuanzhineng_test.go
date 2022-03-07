package old

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_xunhuanzhineng(t *testing.T) {
	//1
	fmt.Println([]int{1, 3, 5, 7}, []int{2, 4, 6, 8, 10, 11, 12})
	//2
	DoMultiTimes(10)
	//3

}

type Input struct { // database model
	Data     []string
	Children []Input
}

type Output struct { // protobuf
	Data     []string
	Children []Output

	DataLength   int // Data长度
	AllSubLength int // 所有子级Children内的Data的length
}

func Transfer(in []Input) (out []Output) {
	if len(in) == 0 {
		return
	}
	for i := range in {
		oItem := Output{
			Data:       in[i].Data,
			Children:   Transfer(in[i].Children),
			DataLength: len(in[i].Data),
		}
		sum := 0
		for j := range oItem.Children {
			if len(oItem.Children[j].Children) == 0 {
				sum += oItem.Children[j].DataLength
			} else {
				sum += oItem.Children[j].AllSubLength
			}
		}
		oItem.AllSubLength = sum
	}
	return
}

func BusyTask(n int) {
	time.Sleep(time.Second)
	println(n)
}

func DoMultiTimes(times int) {
	var wg sync.WaitGroup
	wg.Add(times)
	for i := 1; i <= times; i++ {
		go func(i int) {
			BusyTask(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Merge(a, b []int) (ans []int) {
	m, n := len(a), len(b)
	p1, p2 := 0, 0

	for p1 < m && p2 < n {
		v1, v2 := a[p1], b[p2]
		if v1 <= v2 {
			ans = append(ans, v1)
			p1++
		} else {
			ans = append(ans, v2)
			p2++
		}
	}
	if p1 != m {
		ans = append(ans, a[p1:]...)
	}
	if p2 != n {
		ans = append(ans, b[p2:]...)
	}
	return
}
