// 問題
// 整数 N, K が与えられます。
// N を K 回、改行区切りで出力してください。
package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	for i := 0; i < k; i++ {
		fmt.Println(n)
	}
}
