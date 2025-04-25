// 実数 N が入力されます。N をそのまま出力してください。

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
