package Sort

//Time O(n2),Space O(1)
func SelectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func SelectionSortT(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		maxIndex := i
		for j := i + 1; j < l; j++ {
			if arr[maxIndex] <= arr[j] {
				maxIndex = j
			}
		}
		arr[maxIndex], arr[i] = arr[i], arr[maxIndex]
	}

}
