// 問題
// 整数 N が与えられるので、1 × 2 × ... × (N-1) × N を最大で何回 2 で割ることができるかを求めてください。

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		count += countTwos(i)
	}
	fmt.Println(count)
}

// 何回2で割れるかを返却
func countTwos(n int) int {
	count := 0
	for n%2 == 0 {
		count++
		n /= 2
	}
	return count
}
