// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に、N 個の文字列 s_1, ... , s_N が半角スペース区切りで与えられます。
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
	sc.Scan() // 1行目は読み飛ばす
	inputs := strings.Fields(sc.Text())
	for _, input := range inputs {
		fmt.Println(input)
	}
}
