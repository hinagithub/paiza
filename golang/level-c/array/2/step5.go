// 問題
// 以下の二次元配列を宣言し、この配列の各行の要素数を改行区切りで出力してください。
// 1
// 2 3
// 4 5 6

package main

import "fmt"

func main() {
	matrix := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
	}
	for _, row := range matrix {
		fmt.Println(len(row))
	}
}
