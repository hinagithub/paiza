// 問題
// 3 行 3 列の行列が与えられます。上から i 番目、左から j 番目の整数は a_{i,j} です。
// 3 行 3 列の行列をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 3; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
