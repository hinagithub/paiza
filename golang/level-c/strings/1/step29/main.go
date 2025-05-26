// 問題
// 数値を表す文字列 S , T が与えられるので、S + T の結果を表す文字列を出力してください。
// 繰り上がりが発生する可能性があるので注意してください。

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
