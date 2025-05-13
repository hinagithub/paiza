// 問題
// 以下の二次元配列を宣言し、全要素を各行ずつ半角スペース区切りで出力し、行の終わりで改行してください。
// 6 5 4 3 2 1
// 3 1 8 8 1 3

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	matrix := [][]int{{6, 5, 4, 3, 2, 1}, {3, 1, 8, 8, 1, 3}}
	for _, row := range matrix {

		strRow := make([]string, len(row))
		for i, num := range row {
			strRow[i] = strconv.Itoa(num)
		}
		fmt.Println(strings.Join(strRow, " "))
	}
}
