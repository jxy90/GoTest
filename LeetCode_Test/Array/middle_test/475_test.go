package middle_test

import (
	"fmt"
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"sort"
	"testing"
)

func Test_findRadius(t *testing.T) {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))
	fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
	fmt.Println(findRadius([]int{1, 5}, []int{2}))
}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	n := len(heaters)
	var bSearch func(house int) int
	bSearch = func(house int) int {
		if house <= heaters[0] {
			return heaters[0] - house
		} else if house >= heaters[n-1] {
			return house - heaters[n-1]
		}
		l, r := 0, n-1
		for l <= r {
			mid := l + (r-l)>>1
			if heaters[mid] <= house && heaters[mid+1] > house {
				return CommonUtil.Min(house-heaters[mid], heaters[mid+1]-house)
			} else if heaters[mid] <= house {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}
	ans := 0
	for _, house := range houses {
		r := bSearch(house)
		if r == -1 {
			fmt.Println(house)
		}
		ans = CommonUtil.Max(ans, r)
	}
	return ans
}
