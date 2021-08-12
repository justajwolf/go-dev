package lc704

import (
	"fmt"
	"testing"
)

/*
题目704：二分查找
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
*/
func TestBinarySearch(t *testing.T) {
	fmtBinarySearch([]int{-1, 0, 4, 9, 9, 3, 5, 9, 9, 12}, 9, false)
	fmtBinarySearch([]int{-1, 0, 2, 2, 2, 3, 5, 9, 12}, 2, true)
}
func fmtBinarySearch(nums []int, target int, isSort bool) {
	var (
		desc  string
		left  int
		first int
	)
	if isSort {
		desc = "有序"
		first = BinarySearch(&nums, target, 0, len(nums)-1)
		left = BinarySearch2(&nums, target, 0, len(nums)-1)
	} else {
		desc = "无序"
		first = BinarySearch3(&nums, target)
		left = BinarySearch4(&nums, target)
	}
	fmt.Printf("输入%s：%+v, %d\n", desc, nums, target)
	fmt.Printf("输出1(第一个index)：%d\n", first)
	fmt.Printf("输出2(最左侧index)：%d\n", left)
}
