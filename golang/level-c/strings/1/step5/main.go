// 問題
// 1 行で文字列 S が与えられるので、 S の文字数を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(len(scanner.Text()))
}
