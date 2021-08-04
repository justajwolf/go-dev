package sorts

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	fmtHeapSort([]int{2, 6, 8, 3, 7, 9, 0, 1})
}
func fmtHeapSort(arr []int) {
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出(升序)：%+v\n", HeapSortAsc(arr))
	fmt.Printf("输出(降序)：%+v\n", HeapSortDesc(arr))
}
