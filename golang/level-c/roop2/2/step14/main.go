// 問題
// 整数 N が与えられます。
// N の約数の個数を出力してください。
// 約数とは、N を割り切る整数のことを指します。

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}

	fmt.Println(count)
}
