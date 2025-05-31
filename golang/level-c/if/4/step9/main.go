// 問題
// -1,000 ≦ A ≦ B ≦ 1,000 を満たす 2 つの整数 A, B が与えられます。
// A 以上 B 以下である 2 つの整数 X, Y を適当に選んだとき、X * Y の取り得る値の最小値を出力してください。
// なお、X と Y は同じ値でもよいものとします。
// たとえば A が 3, B が 5 の場合について考えます。
// これは X と Y を両方とも 3 にしたときが最小で、 X * Y が 9 となります。

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	min := int(1e9)
	for x := a; x <= b; x++ {
		for y := a; y <= b; y++ {
			prod := x * y
			if prod < min {
				min = prod
			}
		}
	}
	fmt.Println(min)
}
