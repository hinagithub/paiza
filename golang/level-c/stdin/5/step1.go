// 問題
// 整数 a が与えられるので a を出力してください。

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
