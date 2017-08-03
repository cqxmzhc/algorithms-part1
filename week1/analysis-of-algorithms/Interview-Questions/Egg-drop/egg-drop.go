package main

import "math"

func main() {

}

//n eggs; k floors, 递归
//(n,k) => (n-1, k-1); (n, k-i)
func eggDrop(n, k int) int {
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, k)
	}

	//n eggs, one/zero floor
	for i := 1; i <= n; i++ {
		res[i][0] = 0
		res[i][1] = 1
	}

	// 1/0 egg, n floors
	for j := int(1); j <= k; k++ {
		res[1][j] = j
	}

	for i := int(2); i <= n; i++ {
		for j := int(2); j <= k; j++ {
			//从一楼开始，计算出从任意一层扔时最坏情况下的需要扔的次数， 取最小值
			for x := 1; x <= j; x++ {
				max := 1 + int(math.Max(float64(res[i-1][j-1]), float64(res[i][k-j])))
				if res[i][j] > max {
					res[i][j] = max
				}
			}
		}
	}

	//最坏情况下最少需要扔的次数, 时间复杂度为O(nk^2)
	return res[n][k]
}
