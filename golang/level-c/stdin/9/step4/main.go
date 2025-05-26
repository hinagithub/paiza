// 問題
// 1 行目に整数 N が与えられます。
// 2 行目以降に、N 組の文字列 s_i と整数 a_i が N 行で与えられます。(1 ≦ i ≦ N)
// N 組の s_i と a_i を改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // Nを読み込飛ばす
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
