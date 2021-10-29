package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{randomDoublyNode(10)},
			"10<->9<->8<->7<->6<->5<->4<->3<->2<->1",
		},
		{
			"success",
			args{randomDoublyNode(2)},
			"2<->1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.args.head)
			assert.Equal(t, tt.want, got.String())
		})
	}
}
