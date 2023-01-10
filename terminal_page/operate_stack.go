package main

type OperateStack[T any] struct {
	stack []T
}

func (s *OperateStack[T]) Push(v T) {
	s.stack = append(s.stack, v)
}
func (s *OperateStack[T]) Pop() T {
	if len(s.stack) == 0 {
		panic("stack is empty")
	}
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}
func (s *OperateStack[T]) Top() T {
	if len(s.stack) == 0 {
		panic("stack is empty")
	}
	return s.stack[len(s.stack)-1]
}

func (s *OperateStack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}
