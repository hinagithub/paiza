// 問題
// 整数 N が与えられます。
// 2 の 1 乗から 2 の N 乗までを改行区切りで出力してください。

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	result := 2
	for i := 1; i <= n; i++ {
		fmt.Println(result)
		result *= 2
	}
}
