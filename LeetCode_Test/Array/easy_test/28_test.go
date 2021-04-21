package easy_test

func strStr(haystack string, needle string) int {
	h, n := len(haystack), len(needle)
	i, j := 0, 0
	next := getNext(needle)
	for i < h && j < n {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == n {
		return i - j
	}
	return -1
}

//得到next数组
//第j的位置前后缀相同的位置
func getNext(str string) []int {
	n := len(str)
	next := make([]int, n+1)
	next[0] = -1
	j := 0
	k := -1
	for j < n {
		if k == -1 || str[j] == str[k] {
			j++
			k++
			next[j] = k
		} else {
			if next[k] == -1 || next[k] != next[next[k]] {
				k = next[k]
			} else {
				k = next[next[k]]
			}
		}
	}
	return next
}
