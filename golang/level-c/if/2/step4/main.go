// 問題
// 整数Xが与えられます。
// Xが 10 以上ではない場合はYESを、Xが 10 以上である場合はNOを出力してください。

package main

import "fmt"

func main() {
	const (
		Limit = 10
		Yes   = "YES"
		No    = "NO"
	)

	var x int
	fmt.Scan(&x)

	if !(x >= Limit) {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
