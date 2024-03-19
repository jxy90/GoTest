package _022

import (
	"fmt"
	"strings"
	"testing"
)

func Test_tengxun(t *testing.T) {
	fmt.Println("")
}

//

func basename(str string) string {
	//处理/
	strs := strings.Split(str, "/")
	temp := ""
	if len(strs) == 1 {
		temp = str
	} else {
		temp = strs[len(strs)-1]
	}
	//判断.
	if !strings.Contains(temp, ".") {
		return temp
	}
	//处理.
	ans := ""
	temps := strings.Split(temp, ".")
	for i := 0; i < len(temps)-1; i++ {
		ans += temp + "."
	}
	return ans[:len(ans)-1]
}
