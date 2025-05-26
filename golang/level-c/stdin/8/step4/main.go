// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に、N 個の実数 a_1, ... , a_N が半角スペース区切りで与えられます。
// a_1, ... , a_N を改行区切りでそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // 1行目は読み飛ばす
	sc.Scan()
	inputs := strings.Fields(sc.Text())
	for _, input := range inputs {
		fmt.Println(input)
	}
}
