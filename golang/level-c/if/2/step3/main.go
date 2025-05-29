// 問題
// 2 つの整数A、Bが与えられます。
// A, B の少なくとも一方が 10 以上の場合はYESを、そうではない場合はNOを出力してください。

package main

import "fmt"

func main() {
	const (
		Limit = 10
		Yes   = "YES"
		No    = "NO"
	)
	var a, b int
	fmt.Scan(&a, &b)
	if a >= Limit || b >= Limit {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
