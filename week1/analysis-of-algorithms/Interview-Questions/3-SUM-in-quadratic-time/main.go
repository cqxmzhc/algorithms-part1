package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	m := make(map[int]struct{})
	s := make([]int, 50)
	for i := 0; i < 50; i++ {
		v := random(-100, 100)
		s[i] = v
		m[v] = struct{}{}
	}

	sort.Ints(s)
	sum3WithMap(s, m)

	fmt.Println("---------------------------------------")

	sum3WithoutMap(s)
}

func sum3WithMap(p []int, m map[int]struct{}) {
	pl := len(p)

	for i := 0; i < pl; i++ {
		for j := i + 1; j < pl; j++ {
			third := (p[i] + p[j]) * -1
			if _, ok := m[third]; ok {
				//去掉重复的值
				if p[j] < third {
					fmt.Println(p[i], p[j], third)
				}
			}
		}
	}
}

//https://en.wikipedia.org/wiki/3SUM
func sum3WithoutMap(p []int) {
	pl := len(p)
	for i := 0; i <= pl-3; i++ {
		a := p[i]
		start := i + 1
		end := pl - 1
		for {
			if start >= end {
				break
			}

			// fmt.Println(start, end)
			b := p[start]
			c := p[end]
			if a+b+c == 0 {
				fmt.Println(a, b, c)
				start = start + 1
				end = end - 1
			} else if a+b+c > 0 {
				end = end - 1
			} else if a+b+c < 0 {
				start = start + 1
			}
		}

	}

}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
