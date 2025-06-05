// 問題
// 整数 N が与えられます。
// N の階乗 N! の末尾に 0 がいくつ付くか求め、出力してください。
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	factorial := big.NewInt(1)
	for i := 1; i <= n; i++ {
		bigI := big.NewInt(int64(i))
		factorial.Mul(factorial, bigI)
	}

	strFactorial := factorial.String()
	zeroCount := 0

	// 末尾から走査して0をカウント
	for i := len(strFactorial) - 1; i >= 0; i-- {
		if strFactorial[i] == '0' {
			zeroCount++
		} else {
			break
		}
	}

	fmt.Println(zeroCount)
}

// 高速にやるならこっち
// 末尾の 0 の個数 = N! の中に「2×5」のペアがいくつあるかを調べている
// func countTrailingZeros(n int) int {
// 	count := 0
// 	for n > 0 {
// 		n /= 5
// 		count += n
// 	}
// 	return count
// }
