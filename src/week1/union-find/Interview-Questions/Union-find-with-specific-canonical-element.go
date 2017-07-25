package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	uf := UF{}
	uf.init(1000)

	counter := 0
	for {
		counter = counter + 1
		p := random(0, 1000)
		q := random(0, 1000)

		// fmt.Printf("p:%d q:%d\n", p, q)

		if allConnected := uf.union(p, q); allConnected {
			break
		}
	}

	fmt.Printf("counter:%d\n", counter)

	max := uf.find(455)
	fmt.Printf("Max value: want %d got %d\n", 999, max)
}

type UF struct {
	Friends   []int
	Component []int
	Max       []int
}

func (uf *UF) init(n int) {
	for i := 0; i < n; i++ {
		uf.Friends = append(uf.Friends, i)
		//空间复杂度O(n)
		uf.Component = append(uf.Component, 1)
		//保存每个Component中的最大值
		uf.Max = append(uf.Max, i)
	}
}

func (uf *UF) root(p int) int {
	for {
		if p != uf.Friends[p] {
			uf.Friends[p] = uf.Friends[uf.Friends[p]]
			p = uf.Friends[p]
		} else {
			return p
		}
	}
}

func (uf *UF) connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf *UF) union(p, q int) bool {
	//root表示被合并的树的根
	root := 0
	if !uf.connected(p, q) {
		//O(lgn)
		pRoot := uf.root(p)
		qRoot := uf.root(q)
		if uf.Component[pRoot] >= uf.Component[qRoot] {
			uf.Friends[qRoot] = pRoot

			//保存合并后的最大值
			if uf.Max[pRoot] < uf.Max[qRoot] {
				uf.Max[pRoot] = uf.Max[qRoot]
			}

			//合并后树的大小为两棵树大小之和
			uf.Component[pRoot] = uf.Component[pRoot] + uf.Component[qRoot]
			root = pRoot
		} else {
			uf.Friends[pRoot] = qRoot
			if uf.Max[qRoot] < uf.Max[pRoot] {
				uf.Max[qRoot] = uf.Max[pRoot]
			}
			uf.Component[qRoot] = uf.Component[qRoot] + uf.Component[pRoot]
			root = qRoot
		}
	}

	return uf.allConnected(root)
}

func (uf *UF) allConnected(root int) bool {
	//每次合并后都检查被合并树的大小， 如果大小等于所有元素则认为所有的元素都连接
	//O(1)
	if uf.Component[root] == len(uf.Friends) {
		return true
	}

	return false
}

func (uf *UF) find(find int) int {
	root := uf.root(find)
	return uf.Max[root]
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min

}
