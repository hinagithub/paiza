// 問題
// 以下の二次元配列を宣言し、この配列の行数を出力してください。
// 1 2 3
// 4 5 6
// 8 1 3

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{8, 1, 3},
	}
	fmt.Println(len(matrix))
}
