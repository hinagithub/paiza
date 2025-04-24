// 問題
// 自然数 N が与えられます。
// N が 3 けたになるよう数値の前に 0 を埋めて出力してください。

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
	fmt.Printf("%03d\n", n)
}
