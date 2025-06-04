// 問題
// 現在所持金を A 円持っています。
// 所持金が毎日 10% ずつ増えるとき、何日後に B 円を超えるか出力してください。
// また、増加するお金は小数点以下切り捨てで考えることとします。
// 例として、所持金が 831 円 のとき、10% は 83.1円 ですが、増加するお金は 83 円 です。

package main

import (
	"fmt"
	"math"
)

func main() {
	const taxRate = 0.1
	var pocketMoney, border int
	fmt.Scan(&pocketMoney, &border)

	dayCount := 0

	// borderを越えるまで所持金を10%増やし続ける
	for pocketMoney <= border {
		raise := math.Floor(float64(pocketMoney) * taxRate)
		pocketMoney += int(raise)
		dayCount++
	}
	fmt.Println(dayCount)
}
