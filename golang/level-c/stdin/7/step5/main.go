// 問題
// 1 行目に、整数 N と、続けて N 個の文字列 s_1, ... , s_N が半角スペース区切りで与えられます。
// s_1, ... , s_N を改行区切りで出力してください。

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
	inputs := strings.Fields(sc.Text())
	for _, input := range inputs[1:] {
		fmt.Println(input)
	}
}
