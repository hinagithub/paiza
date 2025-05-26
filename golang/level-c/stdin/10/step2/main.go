// 問題
// 1 行目で整数 N が与えられます。
// 2 行目以降で N 行 3 列の行列が与えられます。上から i 番目、左から j 番目の整数は a_{i,j} です。
// N 行 3 列の行列をそのまま出力してください。

// 入力例
// 2
//1 2 3
// 8 1 3

// 出力例
// 1 2 3
// 8 1 3

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // Nを読み込飛ばす

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
