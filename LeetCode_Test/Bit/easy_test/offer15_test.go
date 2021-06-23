package easy_test

func hammingWeight_offer15(num uint32) int {
	ans := 0
	for num != 0 {
		num = num & (num - 1)
		ans++
	}
	return ans
}
