package easy_test

import (
	"fmt"
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}
func pancakeSort(arr []int) (ans []int) {
	for i := len(arr); i > 1; i-- {
		index := 0
		for j, v := range arr[:i] {
			if v > arr[index] {
				index = j
			}
		}
		if index == i {
			continue
		}
		for j := 0; j < (index+1)/2; j++ {
			arr[j], arr[index-j] = arr[index-j], arr[j]
		}
		ans = append(ans, index+1)
		for j := 0; j < i/2; j++ {
			arr[j], arr[i-j-1] = arr[i-j-1], arr[j]
		}
		ans = append(ans, i)
	}
	return
}
