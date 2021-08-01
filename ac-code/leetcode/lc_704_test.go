package leetcode

import (
	"fmt"
	"testing"
)

/*
题目704：二分查找
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
*/
func TestBinarySearch(t *testing.T) {
	nums, target := []int{-1, 0, 4, 9, 9, 3, 5, 9, 9, 12}, 9
	fmt.Printf("输入无序：%+v, %d\n", nums, target)
	fmt.Printf("输出1(第一个index)：%+v\n", BinarySearch(&nums, target, 0, len(nums)-1))
	fmt.Printf("输出2(最左侧index)：%+v\n", BinarySearch2(&nums, target, 0, len(nums)-1))

	nums, target = []int{-1, 0, 2, 2, 2, 3, 5, 9, 12}, 2
	fmt.Printf("输入有序：%+v, %d\n", nums, target)
	fmt.Printf("输出1(第一个index)：%+v\n", BinarySearch3(&nums, target))
	fmt.Printf("输出2(最左侧index)：%+v\n", BinarySearch4(&nums, target))
}

// 查找数组->递归实现(无序)
func BinarySearch(nums *[]int, target, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if (*nums)[mid] == target {
		return mid
	}
	leftIndex := BinarySearch(nums, target, start, mid-1)
	if leftIndex != -1 {
		return leftIndex
	}
	return BinarySearch(nums, target, mid+1, end)
}

// 查找数组->递归实现(无序, 并且找到第一个出现的target得index)
func BinarySearch2(nums *[]int, target, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if (*nums)[mid] == target {
		// 如果找到一个，继续往左找，直到找不到返回mid，找到就返回最左边得index
		leftIndex := BinarySearch2(nums, target, start, mid-1)
		if leftIndex != -1 {
			return leftIndex
		}
		return mid
	}
	leftIndex := BinarySearch2(nums, target, start, mid-1)
	if leftIndex != -1 {
		return leftIndex
	}
	return BinarySearch2(nums, target, mid+1, end)
}

// 查找数组->循环实现(有序)
func BinarySearch3(nums *[]int, target int) int {
	start, end := 0, len(*nums)-1
	for start <= end {
		mid := (start + end) / 2
		switch {
		case (*nums)[mid] == target:
			return mid
		case (*nums)[mid] > target:
			end = mid - 1
		case (*nums)[mid] < target:
			start = mid + 1
		}
	}
	return -1
}

// 查找数组->循环实现(有序，找到第一个出现的target得index)
func BinarySearch4(nums *[]int, target int) int {
	index, start, end := -1, 0, len(*nums)-1
	for start <= end {
		mid := (start + end) / 2
		switch {
		case (*nums)[mid] == target:
			if index == -1 || mid < index {
				index = mid
			}
			end = mid - 1
		case (*nums)[mid] > target:
			end = mid - 1
		case (*nums)[mid] < target:
			start = mid + 1
		}
	}
	return index
}
