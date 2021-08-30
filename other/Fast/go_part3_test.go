package Fast

import (
	"fmt"
	"math"
	"testing"
)

func Test06(t *testing.T) {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
	fmt.Println(piggyBank == 3)
	fmt.Println(math.Abs(piggyBank-0.3) <= 0.001)
	fmt.Println("----------")
	//1
	//0.30000000000000004
	//可以看到,浮点类型不适合做金融计算
	//为了尽量小化设入错误,建议先乘法后除法
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "° F")
	fmt.Println((celsius*9.0/5.0)+32, "° F")

	fmt.Println("----------")
	fmt.Printf("%T\r\n", 1)

	fmt.Println("----------")
	fmt.Printf("%c", 56)
	fmt.Println("----------")
	fmt.Println('1' - 1)
	fmt.Println("----------")
	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("converted:", v8)
	}
	fmt.Println("----------")
}
