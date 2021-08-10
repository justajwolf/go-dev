package sorts

// 插入排序：从小到大排序
func InsertSortAsc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		input := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > input {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = input
	}
	return arr
}

// 插入排序：从大到小排序
func InsertSortDesc(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		input := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < input {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = input
	}
	return arr
}
