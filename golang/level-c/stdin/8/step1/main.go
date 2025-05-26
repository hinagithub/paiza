// 問題
// 1 行目に整数 N が与えられます。
// 2 行目以降に、N 個の実数 a_1, ... , a_N が N 行で与えられます。
// a_1, ... , a_N を改行区切りでそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // 1行目は読み飛ばす
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
