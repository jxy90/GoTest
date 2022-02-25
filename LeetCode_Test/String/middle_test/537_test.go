package middle_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_complexNumberMultiply(t *testing.T) {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
}

func complexNumberMultiply(num1 string, num2 string) string {
	real1, image1 := complexNumberMultiplyHelper(num1)
	real2, image2 := complexNumberMultiplyHelper(num2)
	return fmt.Sprintf("%v+%vi", real1*real2-image1*image2, real1*image2+real2*image1)
}

func complexNumberMultiplyHelper(num string) (real, image int) {
	index := strings.Index(num, "+")
	real, _ = strconv.Atoi(num[:index])
	image, _ = strconv.Atoi(num[index+1 : len(num)-1])
	return
}
