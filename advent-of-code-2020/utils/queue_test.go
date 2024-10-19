package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	q := NewQueue[int](5)
	require.Equal(t, 0, q.Len())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	require.Equal(t, 5, q.Len())
	q.Push(6)
	require.Equal(t, 5, q.Len())

	list := []int{}
	for iter := q.Iterator(); iter.HasNext(); {
		elem := iter.Next()
		list = append(list, elem)
	}
	require.Equal(t, []int{2, 3, 4, 5, 6}, list)

}