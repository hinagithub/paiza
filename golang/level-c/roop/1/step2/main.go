// 問題
// 正の整数 N が与えられます。
// 1 ~ N の整数を 1 から順に改行区切りで出力してください。

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
