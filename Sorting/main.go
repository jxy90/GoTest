package main

import (
	"fmt"
	"github.com/jxy90/GoTest/Sorting/Sort"
)

func main() {
	arr := []int{99, 23, 21, 6, 7, 54, 2, 84, 51, 12, 8, 71, 48, 1, 29}
	fmt.Println(len(arr))
	Sort.InsertionSort(arr)
	fmt.Println(arr)
}
