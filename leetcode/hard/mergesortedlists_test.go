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
	}{
		{
			name: "test 1",
			lists: []*ListNode{
				{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
				{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				{
					Val: 2,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val:  6,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, MergeKLists(tt.lists))
		})
	}
}
