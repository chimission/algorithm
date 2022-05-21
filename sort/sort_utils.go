package sort

// 比较ab两个切片元素是否相同
func LoopCompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// 复制a切片到新切片里
func MakeCopy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Merge(a, b []int) []int {
	len_a := len(a)
	len_b := len(b)
	len_res := len_a + len_b
	res := make([]int, len_res)
	ai := 0
	bi := 0
	for res_i := 0; ai <= len_a && bi <= len_b && res_i < len_res; res_i += 1 {
		if ai == len_a {
			for ; res_i < len_res; res_i += 1 {
				res[res_i] = b[bi]
				bi += 1
			}
		} else if bi == len_b {
			for ; res_i < len_res; res_i += 1 {
				res[res_i] = a[ai]
				ai += 1
			}
		} else if a[ai] <= b[bi] {
			res[res_i] = a[ai]
			ai += 1
		} else {
			res[res_i] = b[bi]
			bi += 1
		}
	}
	return res
}
