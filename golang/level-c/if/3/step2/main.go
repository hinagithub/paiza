// 整数Nが与えられます。
// Nが 3 の倍数かつ 5 の倍数の場合はYESを、そうではない場合はNOを出力してください。

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
