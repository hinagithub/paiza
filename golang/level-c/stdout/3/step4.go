// 10 個の文字列 S_1, S_2, S_3, ..., S_10 が改行区切りで与えられます。
// これらの文字列をすべて、半角スペース区切りで出力してください。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 10; i++ {
		sc.Scan()
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(sc.Text())
	}
}
