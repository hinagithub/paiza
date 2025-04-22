// 問題
// 自然数 N が入力されます。
// N 行 N 列の九九表を出力してください。具体的には、出力の i 行 j 列目 (1 ≦ i, j ≦ N ) の数値は i * j になるように出力してください。
// ただし、数値の間には半角スペースを、各行の末尾には、半角スペースの代わりに改行を入れてください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(i * j)
			if j == n {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}
