package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_PriorityQueue(t *testing.T) {
	items := map[string]int{
		"AA": 5,
		"BB": 8,
		"CC": 3,
	}
	pq := make(PriorityQueue, len(items))
	i := 0

	for value, priority := range items {
		pq[i] = &Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	item := Item{
		Value:    "DD",
		Priority: 1,
	}
	heap.Push(&pq, &item)
	pq.Update(&item, "EE", 99) // 修改名字并调整优先级

	// EE:99	BB:8	AA:5	CC:3
	for pq.Len() > 0 {
		x := heap.Pop(&pq).(*Item)
		fmt.Printf("%s:%d\t", x.Value, x.Priority)
	}
	fmt.Println()
}
