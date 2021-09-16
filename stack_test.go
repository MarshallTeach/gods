package main

import (
	"strconv"
	"testing"
)

func TestStackPush(t *testing.T) {
	stack := NewStack()
	stackSize := 100

	// Fulfill the test stack
	for i := 0; i < stackSize; i++ {
		value := strconv.Itoa(i)
		stack.Push(value)
	}

	assert(
		t,
		stack.container.Len() == stackSize,
		"stack.container.Len() = %d; want %d", stack.container.Len(), stackSize,
	)

	assert(
		t,
		stack.container.Front().Value == "99",
		"stack.container.Front().Value = %s; want %s", stack.container.Front().Value, "99",
	)

	assert(
		t,
		stack.container.Back().Value == "0",
		"stack.container.Back().Value = %s; want %s", stack.container.Back().Value, "0",
	)
}
