package easy_test

import "testing"

func Test_maximum(t *testing.T) {
	fmt.Println(maximum(1, 2))
}

func maximum(a int, b int) int {
	//整数右移高位补0,负数右移高位补1
	//a>b,ret>0,temp&1=0
	//a<b,ret<0,temp&1=1
	ret := int64(a - b)
	ret = int64(a) - ret&(ret>>63)
	return int(ret)
}
