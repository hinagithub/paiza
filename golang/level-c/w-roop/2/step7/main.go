// 問題
// paiza国では、1 円と X 円と Y 円の 3 種類の硬貨しかありません。
// 支払う硬貨の枚数が最小になるようにz円ちょうどを支払ったときの硬貨の枚数を求めてください。
// ただし、支払う各硬貨の枚数に制限は無いものとします。

package main

import (
	"fmt"
	"math"
)

func main() {
	var X, Y, total int
	fmt.Scan(&X, &Y, &total)

	minCoins := math.MaxInt32

	// X硬貨をi枚、Y硬貨をj枚使う組み合わせを全探索
	for i := 0; i <= total/X; i++ {
		for j := 0; j <= total/Y; j++ {
			sum := i*X + j*Y
			if sum <= total {
				coin1 := total - sum
				totalCoins := i + j + coin1
				if totalCoins < minCoins {
					minCoins = totalCoins
				}
			}
		}
	}

	fmt.Println(minCoins)
}

// 日本円ならこれでいい
// package main

// import "fmt"

// func main() {
// 	var coinX, coinY, total int
// 	fmt.Scan(&coinX, &coinY, &total)

// 	coin1 := 1

// 	coinXcount := 0
// 	coinYcount := 0
// 	coin1count := 0

// 	// 大きい方のコインから消費し、使用コイン総枚数を求める
// 	if coinX > coinY {
// 		total, coinXcount = useCoin(total, coinX)
// 		total, coinYcount = useCoin(total, coinY)
// 	} else {
// 		total, coinYcount = useCoin(total, coinY)
// 		total, coinXcount = useCoin(total, coinX)
// 	}
// 	total, coin1count = useCoin(total, coin1)
// 	fmt.Println(coinXcount + coinYcount + coin1count)
// }

// func useCoin(total int, coin int) (int, int) {
// 	count := 0
// 	for total >= coin {
// 		count++
// 		total -= coin
// 	}
// 	fmt.Println(total, count)
// 	return total, count
// }
