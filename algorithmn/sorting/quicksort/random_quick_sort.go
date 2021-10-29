package quicksort

import (
	"math/rand"
)

// 时间复杂度O(nLog(n))
// 空间复杂度O(1)
// 不稳定排序
// 原地排序
func randomQuickSort(in []int) {
	randQSort(in, 0, len(in)-1)
}

func randQSort(in []int, start, end int) {
	if start >= end {
		return
	}
	i := randPartition(in, start, end)
	randQSort(in, start, i-1)
	randQSort(in, i+1, end)
}

func randPartition(in []int, start, end int) int {
	index := rand.Intn(end-start+1) + start
	in[index], in[end] = in[end], in[index]
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
