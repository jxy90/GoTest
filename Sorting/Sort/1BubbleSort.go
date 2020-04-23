package Sort

//时间复杂度O(n2),空间复杂度O(1)
func BubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func BubbleSortT(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i; j < l; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
