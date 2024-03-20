package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(findKthLargest([]int{8, 3456, 8, 34, 76, 383, 25, 6, 23, 45, 7, 8, 56, 34, 3}, 1))

	//mapval := Map{
	//	Cache:  make(map[string]*entry),
	//	Locker: &sync.RWMutex{},
	//}
	//
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		val := mapval.Read(string(i), time.Second*6)
	//		log.Println("读取值为->", val)
	//	}(i)
	//}
	//
	//time.Sleep(time.Second * 3)
	//for i := 0; i < 10; i++ {
	//	go func(val int) {
	//		mapval.Write(string(val), val)
	//	}(i)
	//}
	//
	//time.Sleep(time.Second * 30)

	//print10()
	//printLetter()
}

func findKthLargest(nums []int, k int) []int {
	helper(nums, k)
	return nums
}

func helper(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums)-1
	mid, i := nums[left], 1
	for left < right {
		if nums[i] > mid {
			nums[i], nums[left] = nums[left], nums[i]
			i++
			left++
		} else {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
	helper(nums[:left], k)
	helper(nums[left+1:], k-left-1)
}

func printLetter() {
	fmt.Println()
	wg := sync.WaitGroup{}
	wg.Add(1)
	letter, number := make(chan struct{}), make(chan struct{})

	go func() {
		i := 1
		for {
			select {
			case <-number:
				print(i)
				i++
				print(i)
				i++
				letter <- struct{}{}
			}
		}
	}()
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wg.Done()
					break
				}
				print(string(i))
				i++
				print(string(i))
				i++
				number <- struct{}{}
			}
		}
	}()

	number <- struct{}{}
	wg.Wait()
}

func print10() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch1 := make(chan int)
	ch2 := make(chan int)
	defer func() {
		close(ch1)
		close(ch2)
		ch1 = nil
		ch2 = nil
	}()
	go func() {
		for {
			select {
			case val := <-ch1:
				if val > 10 {
					wg.Done()
					break
				}
				print(val)
				ch2 <- val + 1
			}
		}
	}()
	go func() {
		for {
			select {
			case val := <-ch2:

				print(val)
				ch1 <- val + 1
			}
		}
	}()

	ch1 <- 1
	wg.Wait()
}
