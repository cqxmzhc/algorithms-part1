package main

import (
	"fmt"
)

func main() {
	d := Deque()
	for i := 0; i < 10; i++ {
		d.addFirst(i)
	}
	fmt.Printf("number of deque:%d\n", d.size())

	for i := 10; i < 20; i++ {
		d.addLast(i)
	}
	fmt.Println("number of deque:", d.size())

	i := d.iterator()
	for v := range i {
		fmt.Printf("%d ", v.value)
		fmt.Println()
	}

	d.removeLast()
	i = d.iterator()
	for v := range i {
		fmt.Printf("%d ", v.value)
		fmt.Println()
	}

	d.removeFirst()
	i = d.iterator()
	for v := range i {
		fmt.Printf("%d ", v.value)
		fmt.Println()
	}

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
		num:   0,
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

	if d.last == nil {
		d.last = n
	}
	if d.first == nil {
		d.first = n
	} else {
		d.first.previous = n
		n.next = d.first
		d.first = n
	}

	d.num += 1
}

func (d *deque) addLast(v int) {
	n := &dequeNode{
		value:    v,
		next:     nil,
		previous: nil,
	}

	if d.first == nil {
		d.first = n
	}

	if d.last == nil {
		d.last = n
	} else {
		d.last.next = n
		n.previous = d.last
		d.last = n
	}

	d.num += 1
}

func (d *deque) removeFirst() {
	if d.first != nil {
		d.first = d.first.next
		if d.first != nil {
			d.first.previous = nil
		}
		d.num -= 1
	}

}

func (d *deque) removeLast() {
	fmt.Println("d.last:", d.last.value)
	fmt.Println("d.last.previous:", d.last.previous.value)
	if d.last != nil {
		d.last = d.last.previous
		if d.last != nil {
			d.last.next = nil
		}
		d.num -= 1
	}
}

func (d *deque) iterator() chan *dequeNode {
	c := make(chan *dequeNode, d.num)

	go func() {
		if d.first == nil {
			close(c)
			return
		}

		n := d.first
		for {
			if n != nil {
				c <- n
				n = n.next
			} else {
				break
			}
		}

		close(c)
	}()

	return c
}
