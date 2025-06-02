// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) が与えられます。
// この数列の和を計算し、出力してください。

package main

import (
	"fmt"
)

func main() {
	var n, sum, a int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum += a
	}
	fmt.Println(sum)
}
