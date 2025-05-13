// 問題
// 10 個の整数 a_i (1 ≤ i ≤ 10) が半角スペース区切りで与えられます。
// この数列の全要素を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strs := strings.Fields(scanner.Text())
		for _, str := range strs {
			fmt.Println(str)
		}
	}

}
