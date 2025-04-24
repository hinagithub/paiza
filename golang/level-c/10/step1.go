// 問題
// 文字列 S, T が与えられます。
// S + T = ST という形式で文字列を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	fmt.Printf("%s + %s = %s\n", s, t, s+t)
}
