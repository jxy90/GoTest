package middle_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_optimalDivision(t *testing.T) {
	fmt.Println(optimalDivision([]int{1000}))
	fmt.Println(optimalDivision([]int{1000, 100}))
	fmt.Println(optimalDivision([]int{1000, 100, 10, 2}))
}

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}
	sb := strings.Builder{}
	for i := range nums {
		sb.WriteString(strconv.Itoa(nums[i]))
		if len(nums)-1 == i {
			break
		}
		sb.WriteByte('/')
	}
	ans := sb.String()
	index := strings.Index(ans, "/")
	ans = ans[:index+1] + "(" + ans[index+1:] + ")"
	return ans
}
