package Fast

import (
	"fmt"
	"testing"
)

type kelvin float64
type c float64

func getk() kelvin {
	return 0
}
func getc() c {
	return 1
}
func Test14(t *testing.T) {
	kk := getk()
	//kk = getc()
	fmt.Println(kk)
	fmt.Println("----------")
	//匿名函数就是没有名字的函数，在 Go 里也称作函数字面值。
	//因为函数字面值需要保留外部作用域的变量引用，所以函数字面值都是闭包的。
	//闭包（closure）就是由于匿名函数封闭并包围作用域中的变量而得名的。
	var k kelvin = 294.0
	sensor := func() kelvin {
		return k
	}
	fmt.Println(sensor())
	k++
	fmt.Println(sensor())
	fmt.Println("----------")
	fmt.Println("----------")
}
