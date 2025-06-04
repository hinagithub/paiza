// 問題
// 整数 N が与えられます。
// N が何回 2 で割れるかを求め、出力してください。

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for n%2 == 0 && n > 0 {
		count++
		n = n / 2
	}
	fmt.Println(count)
}
