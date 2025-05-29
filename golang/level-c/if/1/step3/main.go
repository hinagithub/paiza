// 問題
// 整数 A, B, C が与えられます。式 A × B ≦ C が成立している場合はYESを、
// そうではない場合はNOを出力してください。

package main

import "fmt"

func main() {
	const (
		Yes = "YES"
		No  = "NO"
	)
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	ab := a * b

	if ab <= c {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}

}
