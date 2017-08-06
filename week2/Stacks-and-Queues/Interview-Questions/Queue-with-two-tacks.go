package main

import "fmt"

func main() {
	q := newQueue()
	for i := 1; i <= 5; i++ {
		q.enqueue(i)
	}

	for {
		v := q.dequeue()
		if v == -1 {
			break
		}
		fmt.Println(v)
	}
}

//stack对外接口保持不变, push(v), pop() v
type stack struct {
	values []int
}

func (s *stack) push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) pop() int {
	l := len(s.values)
	if l <= 0 {
		return -1
	}

	r := s.values[l-1]
	s.values = s.values[:l-1]
	return r
}

//queue对外接口保持不变
type queue struct {
	pushStack *stack
	popStack  *stack
}

func (q *queue) enqueue(v int) {
	q.pushStack.push(v)
}

func (q *queue) dequeue() int {
	if len(q.popStack.values) == 0 {
		for i := len(q.pushStack.values) - 1; i >= 0; i-- {
			q.popStack.push(q.pushStack.pop())
		}
	}
	return q.popStack.pop()
}

func newQueue() *queue {
	pushStack := &stack{
		values: []int{},
	}
	popStack := &stack{
		values: []int{},
	}

	return &queue{
		pushStack: pushStack,
		popStack:  popStack,
	}
}
