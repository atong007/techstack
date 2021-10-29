package quicksort

// 原地排序-非稳定排序
// 时间复杂度O(nlog(n))
// 空间复杂度O(1)
func quicksort(in []int) {
	qSort(in, 0, len(in)-1)
}

func qSort(in []int, start, end int) {
	if end <= start {
		return
	}

	i := partition(in, start, end)
	qSort(in, start, i-1)
	qSort(in, i+1, end)
}

// 选择一个pivot
// 把大于pivot放右边，小于pivot放左边
func partition(in []int, start, end int) int {
	pivot := in[end]

	i := start
	for j := start; j < end; j++ {
		if in[j] < pivot {
			if i != j {
				in[i], in[j] = in[j], in[i]
			}
			i++
		}
	}
	in[i], in[end] = in[end], in[i]
	return i
}
