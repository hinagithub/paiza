// 問題
// 1 行目で 1 組の整数 a, b が与えられます。
// a と b をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fmt.Println(sc.Text())
}
