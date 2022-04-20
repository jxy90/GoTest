package middle_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"log"
	"testing"
)

func Test_388(t *testing.T) {
	log.Println(lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"))
}

func lengthLongestPath(input string) (ans int) {
	n := len(input)
	levels := make([]int, n)
	for i := 0; i < n; {
		level := 0
		for ; i < n && input[i] == '\t'; i++ {
			level++
		}

		isFile := false
		length := 0
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		//跳过\n
		i++

		if level > 0 {
			length += levels[level-1] + 1
		}
		if isFile {
			ans = CommonUtil.Max(ans, length)
		} else {
			levels[level] = length
		}
	}
	return ans
}
