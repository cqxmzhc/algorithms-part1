package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*随机生成bitonic数组*/

	// a := make([]int, random(10, 100))
	// mid := 0
	// for {
	// 	tmp := random(1, 100)
	// 	if tmp < len(a) {
	// 		mid = tmp
	// 		break
	// 	}
	// }
	//
	// for k, _ := range a {
	// 	a[k] = random(1, 10000)
	// }
	//
	// sort.Ints(a)
	//
	// r := make([]int, len(a))
	// for i := 0; i < mid; i++ {
	// 	r[i] = a[i]
	// }
	//
	// j := mid
	// for i := len(a) - 1; i >= mid; i-- {
	// 	r[j] = a[i]
	// 	j = j + 1
	// }

	r := []int{
		3, 4, 5, 6, 7, 8, 9, 99, 88, 55, 66, 44, 33, 22, 11,
	}
	fmt.Printf("bitonic array:%+v\n", r)

	// bitonicSearch3lgn(r, 88)
	bitonicSearch2lgn(r, 88)
}

//先通过binary search找到最大值O(lgn), 再分别从两个子数组中查找O(2lgn) => O(3lgn)
func bitonicSearch3lgn(p []int, find int) {
	pl := len(p)

	lo := 0
	hi := pl - 1
	maxIndex := -1

	//lgn
	for {
		mid := (hi + lo) / 2

		max := p[mid]
		left := p[mid-1]
		right := p[mid+1]

		if left < max && max > right {
			maxIndex = mid
			fmt.Printf("maxIndex:%d; max value:%d\n", mid, max)
			break
		} else if left < max && max < right {
			lo = mid + 1
		} else if left > max && max > right {
			hi = mid - 1
		}
	}

	//left lgn
	lo = 0
	hi = maxIndex
	for {
		mid := (hi + lo) / 2

		if find == p[mid] {
			fmt.Printf("find value index:%d\n", mid)
			return
		} else if hi == lo {
			fmt.Println("left no match")
			break
		} else if find < p[mid] {
			hi = mid - 1
		} else if find > p[mid] {
			lo = mid + 1
		}
	}

	//right lgn
	lo = maxIndex
	hi = pl - 1
	for {
		mid := (hi + lo) / 2

		if find == p[mid] {
			fmt.Printf("find value index:%d\n", mid)
			return
		} else if hi == lo {
			fmt.Println("right no match")
			return
		} else if find < p[mid] {
			lo = mid + 1
		} else if find > p[mid] {
			hi = mid - 1
		}
	}
}

func bitonicSearch2lgn(p []int, find int) {
	pl := len(p)

	lo := 0
	hi := pl - 1

	for {
		mid := (hi + lo) / 2

		max := p[mid]
		left := p[mid-1]
		right := p[mid+1]

		if find == max {
			fmt.Printf("find value index:%d;\n")
			break
		} else if lo == hi {
			fmt.Println("-------------")
			fmt.Println("no match")
			break
		} else if left < max && max < right {
			fmt.Println("-------------")
			//最大值在右边
			if find > max {
				lo = mid + 1
			} else if find < max {
				for {
					t := mid + 1
					if binarySearch(t, hi, p, find, false) {
						return
					}
				}

				for {
					t := mid - 1
					binarySearch(lo, t, p, find, true)
				}
			}
		} else if left > max && max > right {
			//最大值在左边
			if find > max {
				hi = mid - 1
			} else if find < max {
				for {
					t := mid - 1
					if binarySearch(lo, t, p, find, true) {
						return
					}
				}

				for {
					t := mid + 1
					binarySearch(t, hi, p, find, false)
				}
			}
		}
	}
}

func binarySearch(lo, hi int, p []int, find int, asc bool) bool {
	mid := (lo + hi) / 2

	for {
		if find == p[mid] {
			fmt.Println("find value index:%d\n", mid)
			return true
		} else if lo == hi {
			fmt.Println("no match")
			return false
		}

		if asc {
			if find > mid {
				lo = mid + 1
			} else if find < mid {
				hi = mid - 1
			}
		} else {
			if find > mid {
				hi = mid - 1
			} else if find < mid {
				lo = mid + 1
			}
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
