// 問題
// 整数Nが与えられます。
// Nが 100 以下の場合はYESを、そうではない場合はNOを出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	if n <= 100 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
