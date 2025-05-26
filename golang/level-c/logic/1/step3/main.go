// 問題
// 0 または 1 の整数 A が与えられます。 NOT A の結果を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	if a == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
