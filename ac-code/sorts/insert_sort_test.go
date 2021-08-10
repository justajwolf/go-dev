package sorts

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	fmtInsertSort([]int{2, 6, 8, 3, 7, 9, 0, 1})
}
func fmtInsertSort(arr []int) {
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出(升序)：%+v\n", InsertSortAsc(arr))
	fmt.Printf("输出(降序)：%+v\n", InsertSortDesc(arr))
}
