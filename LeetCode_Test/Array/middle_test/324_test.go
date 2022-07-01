package middle_test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_324(t *testing.T) {
	wiggleSort([]int{1, 5, 1, 1, 6, 4})
}
func wiggleSort(nums []int) {
	n := len(nums)
	arr := append([]int{}, nums...)
	sort.Ints(arr)
	x := (n + 1) / 2
	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}

func wiggleSort0(nums []int) {
	n := len(nums)
	ans := make([]int, 0, n)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
	//3 0,1,2
	//2 0,1
	for i := 0; i < n/2; i++ {
		first, second := i, n-1-i
		// if n%2==0 {
		//     second = i+n/2
		// }
		ans = append(ans, nums[first])
		ans = append(ans, nums[second])
	}
	if n%2 == 1 {
		ans = append(ans, nums[n/2])
	}
	fmt.Println(ans)
	nums = ans
	fmt.Println(nums)
}
