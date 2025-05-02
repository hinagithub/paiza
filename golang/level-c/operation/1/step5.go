// 問題
// 整数 A = 202、B = 134、C = 107 とします。
// ((A+B)*C)^2を計算した結果を出力してください。

package main

import "fmt"

func main() {
	const (
		a = 202
		b = 134
		c = 107
	)

	sumProduct := (a + b) * c
	squared := sumProduct * sumProduct

	fmt.Println(squared)

}
