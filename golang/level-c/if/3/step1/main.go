// 問題
// 整数Nが与えられます。Nが 3 の倍数の場合はYESを、そうではない場合はNOを出力してください。

package main

import "fmt"

func main() {

	const (
		Yes = "YES"
		No  = "NO"
	)

	var n int
	fmt.Scan(&n)

	if n%3 == 0 {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
