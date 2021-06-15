package main

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	fmt.Println("hello test")
	slice := []string{"1", "2", "3", "4", "5"}
	output("slice", slice)
	slice2 := append(slice, "6", "8", "9", "8", "9", "8", "9")
	output("slice2", slice2)
	slice3 := append(slice2, "7", "8", "9")
	output("slice3", slice3)
	slice3[1] = "10"
	fmt.Println("hello test")
	output("slice", slice)
	output("slice2", slice2)
	output("slice3", slice3)
}

func output(label string, slice []string) {
	fmt.Printf("%v:length %v,capacity %v %v \n", label, len(slice), cap(slice), slice)
}
