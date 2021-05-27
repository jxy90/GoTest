package easy_test

func countBits(n int) []int {
	f := make([]int, n+1)
	heightBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			heightBit = i
		}
		f[i] = f[i-heightBit] + 1
	}
	return f
}

func countBits0(n int) []int {
	f := make([]int, n+1)
	for i := range f {
		if i%2 == 1 {
			f[i] = f[i-1] + 1
		} else {
			f[i] = countBitsHelper(i)
		}
	}
	return f
}

func countBitsHelper(num int) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
