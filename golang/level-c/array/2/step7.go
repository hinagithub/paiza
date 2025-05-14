// 問題
// 以下の二次元配列を宣言し、全要素を各行ずつ半角スペース区切りで出力し、各行の終わりで改行してください。

// 1 3 5 7
// 8 1 3 8

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{8, 1, 3, 8},
	}

	for _, row := range matrix {
		strRow := make([]string, len(row))
		for i, num := range row {
			strRow[i] = strconv.Itoa(num)
		}

		fmt.Println(strings.Join(strRow, " "))
	}

}
