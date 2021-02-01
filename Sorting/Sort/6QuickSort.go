package Sort

/*6、快速排序（Quick Sort）
快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

6.1 算法描述
快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：

从数列中挑出一个元素，称为 “基准”（pivot）；
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。*/

//Time O(n*log2n) Space O(n*log2n)		不稳定
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	head, tail := 0, len(arr)-1
	mid := arr[0]
	i := 1
	for head < tail {
		if mid > arr[i] {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		} else {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}
	}
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])
	return arr
}

func QuickSort2(arr []int, start, end int) {
	if start < end {
		left, right := start, end
		mid := arr[(left+right)/2]
		for left <= right {
			for arr[left] < mid {
				left++
			}
			for arr[right] > mid {
				right--
			}
			if left <= right {
				arr[left], arr[right] = arr[right], arr[left]
				left++
				right--
			}
		}
		if right > start {
			QuickSort2(arr, start, right)
		}
		if left < end {
			QuickSort2(arr, left, end)
		}
	}
}
