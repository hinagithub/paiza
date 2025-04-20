// 文字列Sが与えられます。
// Sがpaizaと一致する場合はYESを、一致しない場合はNOを出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	if text == "paiza" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
