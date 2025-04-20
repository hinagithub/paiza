// 問題
// 2 つの文字列 S, T が入力されます。S, T を改行区切りで出力してください。

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
