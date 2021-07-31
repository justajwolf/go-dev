package sorts

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{2, 6, 8, 3, 7, 9, 0, 1}
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出：%+v\n", HeapSort(arr))
}

// 堆排序：大顶堆(从小到大排序)
func HeapSort(arr []int) []int {
	length := len(arr)
	// 初始化构建大顶堆
	for i := length / 2; i >= 0; i-- {
		AdjustHeap(&arr, i, length)
	}
	// 循环调整堆顶，
	for i := length - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		AdjustHeap(&arr, 0, i)
	}
	return arr
}

func AdjustHeap(arr *[]int, k, length int) {
	old_parent := (*arr)[k]
	for i := k * 2; i < length; i *= 2 {
		if i+1 < length && (*arr)[i] < (*arr)[i+1] {
			i++
		}
		if (*arr)[i] > old_parent {
			(*arr)[k] = (*arr)[i]
			k = i
			continue
		}
		break
	}
	(*arr)[k] = old_parent
}
