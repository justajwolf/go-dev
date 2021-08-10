package sorts

import (
	"fmt"
	"testing"
)

func TestMerge2Sort(t *testing.T) {
	fmtMerge2Sort([]int{2, 6, 8, 3, 7, 9, 0, 1})
}
func fmtMerge2Sort(arr []int) {
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出(升序)：%+v\n", Merge2SortAsc(arr))
	fmt.Printf("输出(降序)：%+v\n", Merge2SortDesc(arr))
}
