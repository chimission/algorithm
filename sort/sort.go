package sort

// 插入排序
func InsertSort(list []int) []int {
	length := len(list)
	if length == 1 {
		return list
	} else {
		for pending_index := 1; pending_index < length; pending_index += 1 {
			pending_node := list[pending_index]
			compare_index := pending_index - 1
			for ; (compare_index >= 0) && (list[compare_index] > pending_node); compare_index = compare_index - 1 {
				list[compare_index+1] = list[compare_index]
			}
			list[compare_index+1] = pending_node

		}
	}
	return list
}

//归并排序
func MergeSort(list []int) []int {
	len_list := len(list)
	if len_list == 1 {
		return list
	}
	middle := len_list / 2
	left_list := MergeSort(list[:middle])
	right_list := MergeSort(list[middle:])
	return Merge(left_list, right_list)

}

//冒泡排序
func BubbleSort(list []int) []int {
	len_list := len(list)
	for i := 0; i < len_list-1; i++ {
		for j := len_list - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return list
}
