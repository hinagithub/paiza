// 問題
// 実数Nが入力されます。Nをそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n := sc.Text()
	fmt.Println(n)
}
