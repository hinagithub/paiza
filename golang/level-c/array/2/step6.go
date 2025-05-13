// 問題
// 以下の二次元配列を宣言し、この配列の 2 行目 3 列目の要素を出力してください。
// 1 2 3
// 8 1 3
// 10 100 1

package main

import "fmt"

func main() {
	const (
		x = 2 - 1
		y = 3 - 1
	)
	matrix := [][]int{
		{1, 2, 3},
		{8, 1, 3},
		{10, 100, 1},
	}
	fmt.Println(matrix[x][y])
}
