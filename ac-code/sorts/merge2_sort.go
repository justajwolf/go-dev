package sorts

// 2路归并排序：从小到大排序
func Merge2SortAsc(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	left := Merge2SortAsc(arr[:mid])
	right := Merge2SortAsc(arr[mid:])
	merge := make([]int, length)
	for l, r, i := 0, 0, 0; i < length; i++ {
		switch {
		case l == mid:
			merge[i] = right[r]
			r++
		case r == length-mid:
			merge[i] = left[l]
			l++
		case left[l] <= right[r]:
			merge[i] = left[l]
			l++
		case left[l] > right[r]:
			merge[i] = right[r]
			r++
		}
	}
	return merge
}

// 2路归并排序：从大到小排序
func Merge2SortDesc(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	left := Merge2SortDesc(arr[:mid])
	right := Merge2SortDesc(arr[mid:])
	merge := make([]int, length)
	for l, r, i := 0, 0, 0; i < length; i++ {
		switch {
		case l == mid:
			merge[i] = right[r]
			r++
		case r == length-mid:
			merge[i] = left[l]
			l++
		case left[l] >= right[r]:
			merge[i] = left[l]
			l++
		case left[l] < right[r]:
			merge[i] = right[r]
			r++
		}
	}
	return merge
}
