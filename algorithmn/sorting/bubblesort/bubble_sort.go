package bubblesort

// 时间复杂度O(n^2)
// 空间复杂度O(1)
// 稳定排序
// 原地排序
func bubbleSort(in []int) {
	for i := len(in) - 1; i >= 0; i-- {
		if !sweep(in, i) {
			return
		}
	}
}

func sweep(in []int, compare int) bool {
	var swap bool
	for i := 1; i <= compare; i++ {
		if in[i] < in[i-1] {
			in[i], in[i-1] = in[i-1], in[i]
			swap = true
		}
	}
	return swap
}
