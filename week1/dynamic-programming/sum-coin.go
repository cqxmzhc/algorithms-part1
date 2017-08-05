package main

import "fmt"

func main() {
	coins := []int{1, 2, 3, 5, 7, 8}
	target := 11
	fmt.Println(sumCoin(coins, target))

}

//https://www.topcoder.com/community/data-science/data-science-tutorials/dynamic-programming-from-novice-to-advanced/
//给定一定数量的不同面值的硬币，求得到和为sum需要的最少硬币数目
//
func sumCoin(v []int, target int) (int, []int) {
	//需要的硬币数
	res := map[int]int{}
	//硬币组成
	m := map[int][]int{}

	//base case
	m[0] = []int{}
	res[0] = 0

	for i := 0; i <= target; i++ {
		for _, j := range v {
			if j <= i && (res[i-j]+1 < res[i] || res[i] == 0) {
				res[i] = res[i-j] + 1
				m[i] = append(m[i-j], j)
			}
		}
	}

	return res[target], m[target]

}
