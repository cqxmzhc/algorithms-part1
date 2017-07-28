package main

import (
	"fmt"

	"github.com/cqxmzhc/week1/union-find/Programming-Assignment-1/Percolation"
)

func main() {
	p := &Percolation.Percolation{}
	ps := Percolation.PercolationStats{
		Percolation: p,
	}

	ps.PercolationStat(200, 100)

	fmt.Println(ps.Mean())
	fmt.Println(ps.StandDev())
}
