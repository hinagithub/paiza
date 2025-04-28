// 問題
// 1 行目に整数 N が与えられます。
// 2 行目から (N + 1) 行目までの先頭に整数 M_i (1 ≦ i ≦ N) が与えられます。それに続いて M_i 個の整数 a_1, ..., a_{M_i} が与えられます。
// 上から i 番目、左から j 番目の整数は a_{i,j} です。
// N 行の a_1, ..., a_M をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // Nを読み込飛ばす
	for sc.Scan() {
		texts := strings.Fields(sc.Text())
		fmt.Println(strings.Join(texts[1:], " "))
	}
}
