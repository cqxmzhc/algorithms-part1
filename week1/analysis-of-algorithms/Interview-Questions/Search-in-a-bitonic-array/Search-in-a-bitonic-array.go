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
	bitonicSearch2lgn(r, 1)
	bitonicSearch2lgn(r, 99)
	bitonicSearch2lgn(r, 3)
	bitonicSearch2lgn(r, 22)
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

		middle := p[mid]
		left := p[mid-1]
		right := p[mid+1]

		if left < middle && middle > right {
			maxIndex = mid
			fmt.Printf("maxIndex:%d; middle value:%d\n", mid, middle)
			break
		} else if left < middle && middle < right {
			lo = mid + 1
		} else if left > middle && middle > right {
			hi = mid - 1
		}
	}

	//left lgn
	lo = 0
	hi = maxIndex
	for {
		mid := (hi + lo) / 2

		if find == p[mid] {
			fmt.Printf("find value %d at index:%d\n", find, mid)
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
			fmt.Printf("find value %d at index:%d\n", find, mid)
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

func bitonicSearch2lgn(p []int, find int) bool {
	pl := len(p)

	lo := 0
	hi := pl - 1

	for {
		mid := (hi + lo) / 2

		middle := p[mid]
		left := p[mid-1]
		right := p[mid+1]

		if find == middle {
			fmt.Printf("find value %d at index:%d\n", find, mid)
			return true
		} else if lo == hi {
			fmt.Println("no match")
			return false
		} else if left < middle && middle < right {
			//最大值在右边
			if find > middle {
				//递归执行bitonic search
				lo = mid + 1
			} else if find < middle {
				//期望值在左右都有可能
				//左边，因为最大值在右边，所以左边的数组为递增的数组，使用binary search,时间复杂度O(lgn)
				//右边，因为期望值小于mid位置的值，且mid值到max的值都大于期望值，max右边的数组是递减的数组，所有可以把mid位置右边的整个数组当作排序递减数组使用binary search,时间复杂度为O(lgn)

				//所以最坏情况是先在左边没有找到，再到右边搜索，时间复杂度为O(2lgn)
				t := mid + 1
				if binarySearch(t, hi, p, find, false) {
					return true
				}

				t = mid - 1
				return binarySearch(lo, t, p, find, true)
			}
		} else if left > middle && middle > right {
			//最大值在左边
			if find > middle {
				hi = mid - 1
			} else if find < middle {
				t := mid - 1
				if binarySearch(lo, t, p, find, true) {
					return true
				}

				t = mid + 1
				return binarySearch(t, hi, p, find, false)
			}
		} else if left < middle && middle > right {
			//取到最大值

			t := mid - 1
			if binarySearch(lo, t, p, find, true) {
				return true
			}

			t = mid + 1
			return binarySearch(t, hi, p, find, false)
		}
	}
}

func binarySearch(lo, hi int, p []int, find int, asc bool) bool {
	for {
		mid := (lo + hi) / 2

		if find == p[mid] {
			fmt.Printf("find value %d at index:%d\n", find, mid)
			return true
		} else if lo == hi {
			fmt.Println("left no match")
			return false
		}

		if asc {
			if find > p[mid] {
				lo = mid + 1
			} else if find < p[mid] {
				hi = mid - 1
			}
		} else {
			if find > p[mid] {
				hi = mid - 1
			} else if find < p[mid] {
				lo = mid + 1
			}
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
