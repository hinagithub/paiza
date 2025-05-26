// 問題
// 文字列 s_1, s_2 が半角スペースで区切られて 1 行で与えられます。
// 各文字列を出力するごとに改行し 2 行で出力してください。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	fmt.Println(inputs[0])
	fmt.Println(inputs[1])

}
