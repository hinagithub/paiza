// 問題
// 自然数 Q が与えられます。Q 回以下の問題に答えてください。

// * 実数 N、自然数 M が入力されます。
// N を丸めて小数第 M 位まで出力してください。
// また、N の小数部が小数第 M 位に満たない場合は 0 で埋めて出力してください。

// なお、小数第 M 位が 5 になることはありません。
// 自然な丸め処理を行って出力すると正解になります。

// 入力例
// 4
// 0.813 1
// 0.813 2
// 0.813 3
// 0.813 4

// 出力例
// 0.8
// 0.81
// 0.813
// 0.8130

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	len, _ := strconv.Atoi(sc.Text())
	for i := 1; i <= len; i++ {
		// 変数設定
		sc.Scan()
		inputs := strings.Split(sc.Text(), " ")
		n, _ := strconv.ParseFloat(inputs[0], 64)
		decimalPlace, _ := strconv.Atoi(inputs[1])

		// 処理と出力
		rounded := roundDecimal(n, decimalPlace)
		format := fmt.Sprintf("%%.%df\n", decimalPlace)
		fmt.Printf(format, rounded)
	}
}

func roundDecimal(n float64, decimalPlace int) float64 {
	multiplier := math.Pow(10, float64(decimalPlace))
	result := n * multiplier
	result = math.Round(result)
	result = result / multiplier
	return result
}
