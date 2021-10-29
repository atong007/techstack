package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_randomDoublyNode(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{10},
			"1<->2<->3<->4<->5<->6<->7<->8<->9<->10",
		},
		{
			"success",
			args{1},
			"1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := randomDoublyNode(tt.args.count)
			assert.Equal(t, tt.want, got.String())
		})
	}
}
