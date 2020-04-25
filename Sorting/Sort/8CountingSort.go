package Sort

//Time O(n+k) 	Space O(n+k)
func CountingSort(arr []int, maxValue int) []int {
	bucket := make([]int, maxValue+1) //Space O(n+k)
	for _, v := range arr {
		bucket[v] += 1
	}
	i := 0
	for k, v := range bucket {
		for v > 0 { //Space O(n+k)
			arr[i] = k
			v--
			i++
		}
	}
	return arr //Space O(n+k)
}
