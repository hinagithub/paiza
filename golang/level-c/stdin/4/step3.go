// 問題
// 整数 a_1, a_2, a_3, a_4, a_5 が 5 行で与えられるので a_1, a_2, a_3, a_4, a_5 を 5 行で出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
