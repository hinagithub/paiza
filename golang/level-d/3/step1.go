// 問題
// 入力された文字列をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	fmt.Println(text)
}
