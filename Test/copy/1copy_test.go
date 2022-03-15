package test

import (
	"sync"
	"testing"
)

/**
* 深拷贝
 */
func TestSliceDeepCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5, 5)
	// 深拷贝
	copy(slice2, slice1)
	t.Log(slice1, len(slice1), cap(slice1))
	// [1 2 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2))
	// [1 2 3 4 5] 5 5
	slice1[1] = 100
	t.Log(slice1, len(slice1), cap(slice1))
	// [1 100 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2))
	// [1 2 3 4 5] 5 5
}

/**
* 切片非并发安全
* 多次执行，每次得到的结果都不一样
* 可以考虑使用 channel 本身的特性 (阻塞) 来实现安全的并发读写
 */
func TestSliceConcurrencySafe(t *testing.T) {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log(len(a))
	// not equal 10000
}

func TestSliceEmptyOrNil(t *testing.T) {
	var slice1 []int
	// slice1 is nil slice
	slice2 := make([]int, 0)
	// slcie2 is empty slice
	var slice3 = make([]int, 2)
	// slice3 is zero slice
	if slice1 == nil {
		t.Log("slice1 is nil.")
		// 会输出这行
	}
	if slice2 == nil {
		t.Log("slice2 is nil.")
		// 不会输出这行
	}
	t.Log(slice3) // [0 0]
}

// 注意：Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
// 拷贝的内容是非引用类型（int、string、struct等这些），在函数中就无法修改原内容数据；
// 拷贝的内容是引用类型（interface、指针、map、slice、chan等这些），这样就可以修改原内容数据。
func TestSliceFn(t *testing.T) {
	// 参数为引用类型slice：外层slice的len/cap不会改变，指向的底层数组会改变
	s := []int{1, 1, 1}
	newS := sliceAppend(s)
	// 函数内发生了扩容
	t.Log(s, len(s), cap(s))
	// [1 1 1] 3 3
	t.Log(newS, len(newS), cap(newS))
	// [1 1 1 100] 4 6

	s2 := make([]int, 0, 5)
	newS = sliceAppend(s2)
	// 函数内未发生扩容
	t.Log(s2, s2[0:5], len(s2), cap(s2))
	// [] [100 0 0 0 0] 0 5
	t.Log(newS, newS[0:5], len(newS), cap(newS))
	// [100] [100 0 0 0 0] 1 5

	// 参数为引用类型slice的指针：外层slice的len/cap会改变，指向的底层数组会改变
	sliceAppendPtr(&s)
	t.Log(s, len(s), cap(s))
	// [1 1 1 100] 4 6
	sliceModify(s)
	t.Log(s, len(s), cap(s))
	// [100 1 1 100] 4 6
}

func sliceModify(s []int) {
	s[0] = 100
}

func sliceAppend(s []int) []int {
	s = append(s, 100)
	return s
}

func sliceAppendPtr(s *[]int) {
	*s = append(*s, 100)
	return
}
