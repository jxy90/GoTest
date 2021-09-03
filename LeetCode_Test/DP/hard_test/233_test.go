package DP_test

import (
	"github.com/jxy90/GoTest/Utils/CommonUtil"
	"testing"
)

func Test_countDigitOne(t *testing.T) {
	fmt.Println(countDigitOne(13))
}

//计算每一位1出现的次数
func countDigitOne(n int) int {
	//i表示n的位数,每次进位
	//从个位开始
	//1234567为例,计算十位出现1的次数
	count := 0
	for i := 1; i <= n; i *= 10 {
		//十位的1被循环的次数
		//12345
		a := n / (i * 10)
		//十位
		//67
		b := n % (i * 10)
		//十位大于1234500,出现1的次数
		//10
		c := CommonUtil.Min(CommonUtil.Max(b-i+1, 0), i)
		//12345 + 10
		count += a*i + c
	}
	return count
}

func countDigitOne0(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	ans := 0
	for i := 1; i <= n; i++ {
		j := i / 10
		count := 0
		if i%10 == 1 {
			count = 1
		}
		f[i] = f[j] + count
		ans += f[i]
	}
	return ans
}
