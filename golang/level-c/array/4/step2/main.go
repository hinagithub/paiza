// 問題
// 整数 A, B が与えられます。
// A と B の値を入れ替えて、半角スペース区切りで出力してください。

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// tmp := a
	// a = b
	// b = tmp

	a, b = b, a

	fmt.Println(a, b)
}
