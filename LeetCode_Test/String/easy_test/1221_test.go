package easy_test

func balancedStringSplit(s string) int {
	countL, countR := 0, 0
	ans := 0
	for _, char := range s {
		if char == 'L' {
			countL++
		}
		if char == 'R' {
			countR++
		}
		if countR == countL {
			ans++
		}
	}
	return ans
}
