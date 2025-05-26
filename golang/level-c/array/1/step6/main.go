// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に、N 個の整数 a_i (1 ≤ i ≤ N) が半角スペース区切りで与えられます。
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
	// 1行目は読み飛ばす
	scanner.Scan()
	scanner.Scan()

	inputs := strings.Fields(scanner.Text())
	for _, input := range inputs {
		fmt.Println(input)
	}

}
