package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKNodes(t *testing.T) {
	type args struct {
		head *Node
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{
				head: randomSinglyNode(10),
				from: 3,
				to:   6,
			},
			"1->2->6->5->4->3->7->8->9->10",
		},
		{
			"success",
			args{
				head: randomSinglyNode(10),
				from: 1,
				to:   6,
			},
			"6->5->4->3->2->1->7->8->9->10",
		},
		{
			"success",
			args{
				head: randomSinglyNode(10),
				from: 1,
				to:   10,
			},
			"10->9->8->7->6->5->4->3->2->1",
		},
		{
			"success",
			args{
				head: randomSinglyNode(10),
				from: 2,
				to:   10,
			},
			"1->10->9->8->7->6->5->4->3->2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseKNodes(tt.args.head, tt.args.from, tt.args.to)
			assert.Equal(t, tt.want, got.String())
		})
	}
}
