package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	require.Equal(t, minStack.GetMin(), -3)

	minStack.Pop()

	require.Equal(t, minStack.Top(), 0)

	require.Equal(t, minStack.GetMin(), -2)
}

func TestMinimumEqualToPreviousMinimum(t *testing.T) {
	minStack := Constructor()

	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)

	require.Equal(t, minStack.GetMin(), 0)

	minStack.Pop()

	require.Equal(t, minStack.GetMin(), 0)
}
