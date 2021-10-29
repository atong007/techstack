package bubblesort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		in []int
	}
	randIn := rand.Perm(100)
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"success",
			args{[]int{1, 4, 2, 5, 3}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"success",
			args{[]int{6, 5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"success",
			args{randIn},
			func() []int {
				sort.Ints(randIn)
				return randIn
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.in)
			assert.Equal(t, tt.want, tt.args.in)
		})
	}
}
