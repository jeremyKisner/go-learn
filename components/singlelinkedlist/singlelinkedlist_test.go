package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList_Append(t *testing.T) {
	t.Run("Append", func(t *testing.T) {
		s := &SingleLinkedList{}
		s.Append("1")
		assert.Equal(t, s.Head.Value, "1")
		assert.Equal(t, s.Tail.Value, "1")
		s.Append(1)
		assert.Equal(t, s.Head.Value, "1")
		assert.Equal(t, s.Tail.Value, 1)
		s.Append(2)
		assert.Equal(t, s.Head.Value, "1")
		assert.Equal(t, s.Tail.Previous.Value, 1)
		assert.Equal(t, s.Tail.Value, 2)
	})
}
