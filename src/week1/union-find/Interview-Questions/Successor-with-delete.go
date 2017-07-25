//https://www.sigmainfy.com/blog/interview-questions-successor-with-delete.html

package main

import "fmt"

func main() {

	s := successor{}
	s.init(100)

	s.remove(3)
	fmt.Println(s.suc(3))
	s.remove(1)
	fmt.Println(s.suc(1))
	s.remove(20)
	s.remove(19)
	s.remove(18)
	fmt.Println(s.suc(18))

}

type successor struct {
	ele []int
}

func (s *successor) init(n int) {
	for i := 0; i < n; i++ {
		s.ele = append(s.ele, i)
	}

}

func (s *successor) suc(x int) int {
	//未被删除的元素的successor是它自己
	for {
		if s.ele[x] != x {
			s.ele[x] = s.ele[s.ele[x]]
			x = s.ele[x]
		} else {
			return x
		}
	}
}

func (s *successor) removed(x int) bool {
	//如果x和x+1具有相同的successor,则认为x已经被删除
	return s.suc(x) == s.suc(x+1)
}

//union-find
func (s *successor) remove(x int) {
	if !s.removed(x) {
		sp := s.suc(x)
		sn := s.suc(x + 1)

		//x的successor只能是大于等于x,不需要判断树的大小
		s.ele[sp] = s.ele[sn]
	}
}
