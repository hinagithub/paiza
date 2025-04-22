// 問題
// 九九表を出力してください。具体的には、出力のi行j列目 (1 ≦ i, j ≦ 9) の数値は i * j になるように出力してください。
// ただし、数値の間には半角スペースを、各行の末尾には、半角スペースの代わりに改行を入れてください。

package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Print(i * j)
			if j == 9 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}
