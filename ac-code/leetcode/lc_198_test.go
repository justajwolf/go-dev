package leetcode

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	arr := []int{1, 2, 3, 1}
	fmt.Printf("输入：%+v\n输出：%d\n", arr, Rob(arr))
	arr = []int{2, 7, 9, 3, 1}
	fmt.Printf("输入：%+v\n输出：%d\n", arr, Rob(arr))
}

func Rob(nums []int) int {
	return 1
}
