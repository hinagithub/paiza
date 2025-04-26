// 問題
// 整数 a_1, a_2, ... , a_999, a_1000 が 1,000 行で与えられるので a_1, a_2, ... , a_999, a_1000 を 1,000 行で出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 1000; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
