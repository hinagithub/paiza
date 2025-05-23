// 問題
// 1 行目に整数 N が与えられます。
// 2 行目以降に N 個の点の座標 (x_1, y_1), (x_2, y_2), ..., (x_N, y_N) が与えられます。
// 点 (2, 3) と各点の距離を改行区切りで出力してください。

// 距離の計算にはマンハッタン距離

// |x1 - x2| + |y1 - y2|

// を用いることとします。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const (
		baseX = 2
		baseY = 3
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行目は使わないのでスキップ

	for scanner.Scan() {
		inputs := strings.Fields(scanner.Text())
		if len(inputs) != 2 {
			fmt.Println("input length must be 2")
			os.Exit(1)
		}
		x := a2i(inputs[0])
		y := a2i(inputs[1])
		fmt.Println(diff(x, baseX) + diff(y, baseY))
	}
}

func a2i(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err: %d", err)
		os.Exit(1)
	}
	return num
}

func diff(num int, base int) int {
	diff := base - num
	if diff < 0 {
		return -diff
	}
	return diff
}
