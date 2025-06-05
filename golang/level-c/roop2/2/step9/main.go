// 問題
// 整数 N が与えられます。
// N の階乗 N! を計算して出力してください。

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	fmt.Println(n)

	result := 1
	for i := 1; i < n; i++ {
		result += result * i
	}
	fmt.Println(result)
}
