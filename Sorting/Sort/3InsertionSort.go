package Sort

//Time O(n2),Space O(1)
func InsertionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		preIndex := i + 1
		current := arr[preIndex]
		for ; preIndex > 0 && arr[preIndex-1] > current; preIndex-- {
			arr[preIndex] = arr[preIndex-1]
		}
		arr[preIndex] = current
	}
}

func InsertionSortT(arr []int) {
	l := len(arr)
	for i := 1; i < len(arr); i++ {

	}
}
