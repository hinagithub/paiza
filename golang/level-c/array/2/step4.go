// 問題
// 以下の二次元配列を宣言し、この配列の列数を出力してください。
// 1 2 3 4
// 6 5 4 3
// 3 1 8 1

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{6, 5, 4, 3},
		{3, 1, 8, 1},
	}
	fmt.Println(len(matrix[0]))
}
