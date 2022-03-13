package middle_test

import (
	"fmt"
	"testing"
)

func Test_393(t *testing.T) {
	//fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{235, 140, 4}))
}

func validUtf8(data []int) bool {
	n := len(data)
	if n == 0 {
		return false
	}
	//求个数
	cnt := 0
	temp := data[0]
	for temp > 0 {
		if temp&1 == 1 {
			cnt++
		} else {
			cnt = 0
		}
		temp = temp >> 1
	}
	//个数是否对
	if cnt > n {
		return false
	}
	for i := 1; i < cnt; i++ {
		temp = data[i]
		c := 0
		for temp > 0 {
			if temp^1 == 1 {
				c++
			} else {
				c = 0
			}
			temp = temp >> 1
		}
		if c != 1 {
			return false
		}
	}
	if n > cnt {
		temp = data[cnt]
		c := 0
		for temp > 0 {
			if temp^1 == 1 {
				c++
			} else {
				c = 0
			}
			temp = temp >> 1
		}
		if c == 1 {
			return false
		}
	}
	return true
}
