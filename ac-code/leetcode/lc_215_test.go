package leetcode

import (
	"fmt"
	"testing"
)

/*
题目215：数组中的第K个最大元素
	给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
	请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/
func TestFindKthLargest(t *testing.T) {
	fmtFindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmtFindKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
}
func fmtFindKthLargest(arr []int, k int) {
	fmt.Printf("输入：%+v, %d\n", arr, k)
	fmt.Printf("输出：%+v\n", FindKthLargest(arr, k))
}
