// 問題
// 整数 a, b が 2 行で与えられるので a, b を 2 行で出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
