// 問題
// 文字列 s_1, s_2, s_3 が半角スペースで区切られて 1 行で与えられます。
// 各文字列を出力するごとに改行し 3 行で出力してください。

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
	for _, input := range inputs {
		fmt.Println(input)
	}
}
