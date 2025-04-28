// 問題
// 1 行目で、整数 N と、続けて N 個の整数 a_1, ... , a_N が半角スペース区切りで与えられます。
// a_1, ... , a_N を改行区切りで出力してください。

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
