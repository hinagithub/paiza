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
	for i := 0; i < n; i++ {
		square := 2
		for j := 0; j < i; j++ {
			square *= 2
		}
		fmt.Println(square)
	}
}
