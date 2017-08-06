package main

func main() {
}

type dequeNode struct {
	value    int
	next     *dequeNode
	previous *dequeNode
}

type deque struct {
	num   int
	first *dequeNode
	last  *dequeNode
}

func Deque() *deque {
	return &deque{
		size:  0,
		first: nil,
		last:  nil,
	}
}

func (d *deque) isEmpty() bool {
	return d.num == 0
}

func (d *deque) size() int {
	return d.num
}

func (d *deque) addFirst(v int) {
	n := &dequeNode{
		value:    v,
		next:     nil,
		previous: nil,
	}

	n.next = d.first
	d.first = n
	d.num += 1
}

func (d *deque) addLast(v int) {
	n := &dequeNode{
		value:    v,
		next:     nil,
		previous: nil,
	}

	if d.last == nil {
		d.last = n
	} else {
		d.last.next = n
	}

	d.num += 1
}

func (d *deque) removeFirst() {
	if d.first != nil {
		d.first = d.first.next
		d.num -= 1
	}

}

func (d *deque) removeLast() {
	if d.last != nil {
		d.last = d.last.previous
		d.num -= 1
	}
}

func (d *deque) iterator() chan *dequeNode {

}
