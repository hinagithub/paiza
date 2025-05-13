// 問題
// 数値を表す文字列 S , T が与えられるので、S + T の結果を表す文字列を出力してください。
// なお、 S + T では、 足し算をしたときに全ての桁において繰り上がりが発生しないことが保証されているものとします。

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	a, _ := new(big.Int).SetString(s, 10)
	b, _ := new(big.Int).SetString(t, 10)

	fmt.Println(new(big.Int).Add(a, b))
}
