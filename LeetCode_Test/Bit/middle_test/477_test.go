package middle_test_test

func totalHammingDistance(nums []int) int {
	ans := 0
	for i := 0; i < 30; i++ {
		q := 0
		p := 0
		for j := range nums {
			if nums[j]>>i&1 == 1 {
				q++
			} else {
				p++
			}
		}
		ans += q * p
	}
	return ans
}

func totalHammingDistance0(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			ans += hammingDistance(nums[i], nums[j])

		}
	}
	return ans
}

func hammingDistance(x, y int) int {
	temp := x ^ y
	count := 0
	for temp != 0 {
		if temp&1 == 1 {
			count++
		}
		temp >>= 1
	}
	return count
}
