// 問題
// 文字列 s, t, u が 3 行で与えられるので、s, t, u の 3 行をそのまま出力してください。

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
