// 問題
// 1 行目に整数 N が与えられます。
// 2 行目以降に、N 組の文字列 s_i と整数 a_i が N 行で与えられます。
// 8 組目の s_i と a_i を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const targetIndex = 8
	sc := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for sc.Scan() {
		inputs = append(inputs, sc.Text())
	}
	fmt.Println(inputs[targetIndex])
}
