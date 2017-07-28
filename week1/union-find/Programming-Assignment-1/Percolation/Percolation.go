package Percolation

import (
	"math/rand"
	"time"
)

type Percolation struct {
	virtualTop    *site
	virtualBottom *site
	openSiteNum   int
	lattice       [][]*site
}

type site struct {
	row int
	col int

	parent  *site
	open    bool
	numNode int
	top     bool
}

func (p *Percolation) init(n int) {
	//初始化虚拟节点
	p.virtualTop = &site{
		row:     0,
		col:     0,
		open:    false,
		numNode: 1,
	}
	p.virtualTop.parent = p.virtualTop

	p.virtualBottom = &site{
		row:     0,
		col:     0,
		open:    false,
		numNode: 1,
	}
	p.virtualBottom.parent = p.virtualBottom

	row := make([][]*site, n+1)
	for i := range row {
		for j := 0; j <= n; j++ {
			s := &site{
				row:     i,
				col:     j,
				open:    false,
				numNode: 1,
			}

			//root节点的父亲是它自己
			s.parent = s

			//第一行的为top节点
			if i == 1 {
				s.top = true
			}
			row[i] = append(row[i], s)
		}
	}

	p.lattice = row
}

func (p *Percolation) open(row, col int) {
	if p.indexOutOfRange(row, col) {
		// log.Fatal(fmt.Errorf("index out of range"))
		return
	}

	if !p.isOpen(row, col) {
		p.lattice[row][col].open = true
		p.openSiteNum += 1
	} else {
		return
	}

	s := p.lattice[row][col]
	//联通四个方向的site

	//up
	if p.isOpen(row-1, col) {
		m := p.lattice[row-1][col]
		p.union(s, m)
	}

	//down
	if p.isOpen(row+1, col) {
		m := p.lattice[row+1][col]
		p.union(s, m)
	}

	//left
	if p.isOpen(row, col-1) {
		m := p.lattice[row][col-1]
		p.union(s, m)
	}

	//right
	if p.isOpen(row, col+1) {
		m := p.lattice[row][col+1]
		p.union(s, m)
	}

	//连接顶部或底部的虚拟节点
	if s.row == 1 {
		p.union(s, p.virtualTop)
	} else if s.row == len(p.lattice)-1 {
		p.union(s, p.virtualBottom)
	}
}

func (p *Percolation) isOpen(row, col int) bool {
	if p.indexOutOfRange(row, col) {
		// log.Fatal(fmt.Errorf("index out of range"))
		return false
	}

	return p.lattice[row][col].open
}

func (p *Percolation) isFull(row, col int) bool {
	if p.indexOutOfRange(row, col) {
		// log.Fatal(fmt.Errorf("index out of range"))
		return false
	}
	s := p.lattice[row][col]
	return p.root(s).top
}

func (p *Percolation) NumberOfOpenSite() int {
	return p.openSiteNum
}

func (p *Percolation) percolates() bool {
	//如果顶部和底部的虚拟节点联通，说明系统percolate
	return p.connected(p.virtualTop, p.virtualBottom)
}

func (p *Percolation) indexOutOfRange(row, col int) bool {
	max := len(p.lattice) - 1
	if row <= 0 || col <= 0 || row > max || col > max {
		return true
	}

	return false
}

func (p *Percolation) getSiteValue(row, col int) int {
	return (row-1)*len(p.lattice) + col
}

/*****union methods*****/
func (p *Percolation) union(q, m *site) {
	if !p.connected(q, m) {
		qRoot := p.root(q)
		mRoot := p.root(m)

		if qRoot.numNode > mRoot.numNode {
			mRoot.parent = qRoot
			qRoot.numNode += mRoot.numNode
		} else {
			qRoot.parent = mRoot
			mRoot.numNode += qRoot.numNode
		}
	}
}

func (p *Percolation) connected(q, m *site) bool {
	// return q.parent != nil && (q.parent == m || m.parent == q)
	return p.root(q) == p.root(m)
}

func (p *Percolation) root(s *site) *site {
	for {
		if s.parent != s {
			s.parent = s.parent.parent
			s = s.parent
		} else {
			return s
		}
	}

}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
