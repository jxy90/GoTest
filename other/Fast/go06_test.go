package Fast

import (
	"fmt"
	"testing"
)

func Test06(t *testing.T) {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
	//1
	//0.30000000000000004
	//可以看到,浮点类型不适合做金融计算
	//为了尽量小化设入错误,建议先乘法后除法
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "° F")
	fmt.Println((celsius*9.0/5.0)+32, "° F")
}
