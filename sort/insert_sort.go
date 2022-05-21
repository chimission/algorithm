package sort

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
