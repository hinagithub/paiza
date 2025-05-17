// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数の順番を反転させ、改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行目は使わないので飛ばす
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	// results := make([]string, len(inputs))
	for i := 0; i < len(inputs); i++ {
		fmt.Println(inputs[len(inputs)-i-1])
	}
}
