// 問題
// s_1, s_2, s_3, ... s_999, s_1000 の 1,000 個の文字列が与えられます。
// 文字列を与えられた順番通りに出力してください。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const len int = 1000
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < len; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
