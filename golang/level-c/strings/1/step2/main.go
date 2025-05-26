// 問題
// 1 行で文字列 S が与えられるので、その文字列をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}
