// 問題
// 整数 N, M, K が与えられます。
// N が M ずつ増えるとき、何回目に K を越えるか出力してください。

package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)

	count := 0

	// kを越えるまでn+mし続ける
	for n <= k {
		count++
		n += m
	}
	fmt.Println(count)
}
