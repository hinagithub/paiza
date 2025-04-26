// 問題
// 整数 a_1, a_2, ... , a_99, a_100 が 100 行で与えられるので a_1, a_2, ... , a_99, a_100 を 100 行で出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 100; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
