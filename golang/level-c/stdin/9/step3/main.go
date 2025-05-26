// 問題
// 1 行目に整数 N が与えられます。
// 2 行目以降に、N 組の整数 a_i と b_i が N 行で与えられます。(1 ≦ i ≦ N)
// 8 組目の a_i と b_i を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for sc.Scan() {
		inputs = append(inputs, sc.Text())
	}
	fmt.Println(inputs[8])
}
