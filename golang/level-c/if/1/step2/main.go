// 問題
// 整数Nが与えられます。Nが 100 以下の場合はYESを、そうではない場合はNOを出力してください。

package main

import "fmt"

func main() {
	const (
		Limit = 100
		Yes   = "YES"
		No    = "NO"
	)

	var num int
	fmt.Scan(&num)

	if num <= Limit {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}
}
