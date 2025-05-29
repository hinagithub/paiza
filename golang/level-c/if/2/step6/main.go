// 問題
// 3 つの整数X, Y, Zが与えられます。「Xが 10 以上」かつ「Yが 10 以上」の場合はYESを、
// そうではない場合はNOを出力してください。
// ただし、「Zが 10 以上の」場合はXとYの値にかかわらず、必ずYESを出力してください。

package main

import "fmt"

func main() {
	const (
		Limit = 10
		Yes   = "YES"
		No    = "NO"
	)
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if z >= 10 || (x >= 10 && y >= 10) {
		fmt.Println(Yes)
	} else {
		fmt.Println(No)
	}

}
