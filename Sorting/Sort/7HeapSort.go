package Sort

/*7、堆排序（Heap Sort）
堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

7.1 算法描述
将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。*/
var length int

func buildMaxHeap(arr []int) {
	length = len(arr)
	for i := length / 2; i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapify(arr []int, i int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < length && arr[left] > arr[largest] {
		largest = left
	}
	if right < length && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, largest)
	}

}

//Time O(n*log2n) Time O(1) 	不稳定
func HeapSort(arr []int) []int {
	buildMaxHeap(arr)
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		length -= 1
		heapify(arr, 0)
	}
	return arr
}
