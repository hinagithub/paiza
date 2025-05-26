// 問題
// 文字列 s が 1 行で与えられます。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text1 := sc.Text()
	sc.Scan()
	text2 := sc.Text()
	fmt.Println(text1)
	fmt.Println(text2)
}
