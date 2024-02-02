package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    []*ListNode
		expected *ListNode
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, MergeKLists(tt.lists))
		})
	}
}
