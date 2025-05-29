// 問題
// 整数 N が与えられます。
// Nが 0 ではない場合はYESを、 0 である場合はNOを出力してください。

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
