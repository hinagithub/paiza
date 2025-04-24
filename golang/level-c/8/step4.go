// 問題
// 実数 N、自然数 M が入力されます。N を丸めて小数第 M 位まで出力してください。
// また、N の小数部が小数第 M 位に満たない場合は 0 で埋めて出力してください。

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
	inputs := strings.Split(sc.Text(), " ")
	n, _ := strconv.ParseFloat(inputs[0], 64)
	decimalPlace, _ := strconv.Atoi(inputs[1])
	rounded := round(n, decimalPlace)
	format := fmt.Sprintf("%%.%df", decimalPlace)
	fmt.Printf(format, rounded)
}

func round(n float64, decimalPlace int) float64 {
	multiplier := math.Pow(10, float64(decimalPlace))
	result := n * multiplier
	result = math.Round(result)
	result = result / multiplier
	return result
}
