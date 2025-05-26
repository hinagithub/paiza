// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数に含まれている M の個数を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	target := strings.Fields(scanner.Text())[1]
	scanner.Scan()

	count := 0
	for _, str := range strings.Fields(scanner.Text()) {
		if str == target {
			count++
		}
	}

	fmt.Println(count)
}
