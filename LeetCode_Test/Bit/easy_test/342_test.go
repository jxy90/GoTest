package easy_test

func isPowerOfFour(n int) bool {
	temp := 1
	for temp < n {
		temp *= 4
	}
	return temp == n
}
func isPowerOfFour0(n int) bool {
	return n > 0 && n&(n-1) == 0 && n%3 == 1
}
