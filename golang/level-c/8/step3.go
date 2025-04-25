// 問題
// 実数 N が入力されます。N を丸めて小数第 3 位まで出力してください。
// また、N の小数部が小数第 3 位に満たない場合は 0 で埋めて出力してください。
// なお、小数第 4 位が 5 になることはありません。
// 自然な丸め処理を行って出力すると正解になります。

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const decimalPlace int = 3
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.ParseFloat(sc.Text(), 64)
	rounded := roundTo(n)
	format := fmt.Sprintf("%%.%df", decimalPlace)
	fmt.Printf(format, rounded)
}
func roundTo(n float64) float64 {
	return math.Round(n*1000) / 1000
}
