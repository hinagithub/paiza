// 問題
// 2 つの整数A、Bが与えられます。
// AとBが両方とも 10 以上の場合はYESを、そうではない場合はNOを出力してください。

package main

import "fmt"

func main() {
	const (
		limit = 10
		Yes   = "YES"
		No    = "NO"
	)
	var a, b int
	fmt.Scan(&a, &b)

	if a >= limit && b >= limit {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
