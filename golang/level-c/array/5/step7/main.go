// 問題
// 1 行目に整数 N, A, B が与えられます。
// 2 行目以降に N 個の点の座標 (x_1, y_1), (x_2, y_2), ..., (x_N, y_N) が与えられます。
// A 番目の点 と B 番目の点の距離を出力してください。

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

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	baseInfos := strings.Fields(scanner.Text())

	pointA := a2I(baseInfos[1])
	pointB := a2I(baseInfos[2])

	points := [][]int{}

	for scanner.Scan() {
		inputs := strings.Fields(scanner.Text())
		if len(inputs) != 2 {
			fmt.Println("input length must be 2")
			os.Exit(1)
		}
		x := a2I(inputs[0])
		y := a2I(inputs[1])
		points = append(points, []int{x, y})
	}

	resultX := points[pointA-1][0] - points[pointB-1][0]
	resultY := points[pointA-1][1] - points[pointB-1][1]
	fmt.Println(abs(resultX) + abs(resultY))
}

func a2I(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err: %v", err)
		os.Exit(1)
	}
	return num
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
