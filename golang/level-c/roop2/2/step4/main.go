// 問題
// 整数 N, M が与えられます。
// N が何回 M で割れるかを求め、出力してください。

package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	count := 0
	for n%m == 0 && n > 0 {
		count++
		n = n / m
	}
	fmt.Println(count)
}
