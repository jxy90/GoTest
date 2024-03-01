package Array

import (
	"fmt"
	"testing"
)

//二分搜索怎么用？我又总结了套路

//爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, v := range piles {
		if right < v {
			right = v
		}
	}
	for left < right {
		mid := left + (right-left)>>1
		if minEatingSpeedValid(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func minEatingSpeedValid(piles []int, k, h int) bool {
	cost := 0
	for _, pile := range piles {
		cost += pile / k
		if pile%k != 0 {
			cost += 1
		}
	}
	return cost <= h
}

// 在 D 天内送达包裹的能力
func shipWithinDays(weights []int, days int) int {
	sum := 0
	max := 0
	for _, v := range weights {
		sum += v
		if max < v {
			max = v
		}
	}
	left, right := max, sum

	for left < right {
		mid := left + (right-left)>>1
		if shipWithinDaysValid(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func shipWithinDaysValid(weights []int, w, days int) bool {
	cost := 0
	cnt := 1
	for _, v := range weights {
		if cost+v > w {
			cnt++
			cost = v
		} else {
			cost += v
		}
	}
	return cnt <= days
}

func Test_7(t *testing.T) {
	//fmt.Println(minEatingSpeedValid([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{312884470}, 968709470))
}
