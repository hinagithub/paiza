// 問題
// 半角スペースで区切られた文字列 s_1, s_2 が 1 行で与えられます。
// 入力された文字列 s_1, s_2 をそれぞれ改行区切りで出力してください。

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
	text := sc.Text()
	inputs := strings.Split(text, " ")
	fmt.Println(inputs[0])
	fmt.Println(inputs[1])
}
