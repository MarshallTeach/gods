package gods

type Stack struct {
	*Deque
}

func NewStack() *Stack {
	return &Stack{
		Deque: NewDeque(),
	}
}

// Push adds on an item on the top of the stack
func (s *Stack) Push(item interface{}) {
	s.Prepend(item)
}

// Pop removes and returns the item on the top of the stack
func (s *Stack) Pop() interface{} {
	return s.Shift()
}

// Head returns the item on the top of the stack
func (s *Stack) Head() interface{} {
	return s.First()
}
