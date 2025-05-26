// 問題
// 整数 a_1, a_2, a_3, a_4, a_5 が半角スペース区切りで与えられるので、改行区切りにして 5 行で出力してください。

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
