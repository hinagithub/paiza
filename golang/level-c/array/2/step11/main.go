// 問題
// 整数 K と 整数 L が与えられます。
// 以下の二次元配列 a を宣言し、この配列の K 行目 L 列目の要素を出力してください。
// 上から i 番目、左から j 番目の整数は a_ij です。
// 1 2 3 4
// 10 100 0 5
// 8 1 3 8
// 15 34 94 25

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{10, 100, 0, 5},
		{8, 1, 3, 8},
		{15, 34, 94, 25},
	}
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())

	x, _ := strconv.Atoi(inputs[0])
	y, _ := strconv.Atoi(inputs[1])

	fmt.Println(matrix[x-1][y-1])
}
