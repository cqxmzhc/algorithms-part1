package main

import "fmt"

type test struct {
	t []int
}

func main() {
	t1 := test{}
	t1.t = append(t1.t, 1)

	fmt.Println(t1.t)
	passByValue(t1)
	fmt.Println(t1.t)

}

func passByValue(p test) {
	fmt.Println(p.t)
	p.t[0] = 1111
}
