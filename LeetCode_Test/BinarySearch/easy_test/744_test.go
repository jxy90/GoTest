package easy_test

import (
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	//判断数组最大值也小于target的情况
	if letters[r] <= target {
		return letters[0]
	}
	for l < r {
		//l=mid,mid不+1
		mid := l + (r-l)>>1
		if letters[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return letters[l]
}
