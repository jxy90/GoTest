package main

func main() {
	//runtime.GOMAXPROCS(3)

}

const mod = 1337

func superPow(a int, b []int) int {
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		res = res * fPow(a, b[i]) % mod
		a = fPow(a, 10)
	}
	return res
}

func fPow(x, n int) int {
	ans := 1
	for n > 0 {
		if n%2 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return ans
}
func fpow(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 1 {
		return fpow(x, n-1) * x
	} else {
		temp := fpow(x, n/2)
		return temp * temp
	}
}
