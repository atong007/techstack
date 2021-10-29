package singly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWithTraverse(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{randomSinglyNode(10)},
			"10->9->8->7->6->5->4->3->2->1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.args.node.String()
			got := ReverseWithTraverse(tt.args.node)
			assert.Equal(t, tt.want, got.String())
			got = ReverseWithTraverse(got)
			assert.Equal(t, str, got.String())
		})
	}
}

func TestReverseWithRecursive(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{randomSinglyNode(10)},
			"10->9->8->7->6->5->4->3->2->1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.args.node.String()
			got := ReverseWithRecursive(tt.args.node)
			assert.Equal(t, tt.want, got.String())
			got = ReverseWithRecursive(got)
			assert.Equal(t, str, got.String())
		})
	}
}
