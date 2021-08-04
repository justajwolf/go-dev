package leetcode

import "github.com/changbaihe/go-manual/ac-code/sorts"

/*
题目215：数组中的第K个最大元素
	给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
	请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/

func FindKthLargest(nums []int, k int) int {
	length := len(nums)
	// 初始化构建大顶堆
	for i := length / 2; i >= 0; i-- {
		sorts.HeapAdjustMax(&nums, i, length)
	}
	// 循环调整堆顶k-1次，
	for i := length - 1; i > length-1-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		sorts.HeapAdjustMax(&nums, 0, i)
	}
	return nums[length-k]
}
