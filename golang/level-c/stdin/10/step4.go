// 問題
// 1 行目で整数 M が与えられます。
// 2 行目以降で 3 行 M 列の行列が与えられます。上から i 番目、左から j 番目の整数は a_{i,j} です。
// 3 行 M 列の行列をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // Mを読み込飛ばす
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

}
