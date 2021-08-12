package lc53

import (
	"fmt"
	"testing"
)

/*
题目53：最大子序和
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/
func TestMaxSubArray(t *testing.T) {
	fmtMaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmtMaxSubArray([]int{5, 4, -1, 7, 8})
	fmtMaxSubArray([]int{1, 2})
}
func fmtMaxSubArray(arr []int) {
	fmt.Printf("输入：%+v\n", arr)
	fmt.Printf("输出：%+v\n", MaxSubArray(arr))
}
