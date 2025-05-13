// 問題
// 以下の二次元配列を宣言し、要素数を出力してください。
// 1 2 3 4 5 6
// 8 1 3 3 1 8

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4, 5, 6}, {8, 1, 3, 3, 1, 8}}
	count := 0
	for _, line := range matrix {
		count += len(line)
	}
	fmt.Println(count)
}
