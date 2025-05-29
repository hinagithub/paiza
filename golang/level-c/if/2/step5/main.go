// 問題
// 2 つの整数A、Bが与えられます。
// 以下の条件を満たす場合はYESを、そうではない場合はNOを出力してください。
// ・「Aが 10 以上」 かつ 「Bが 10 以上ではない」

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

	if a >= Limit && !(b >= Limit) {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
