// 問題
// 自然数 N が与えられます。
// N が 3 けたになるよう数値の前に半角スペースを埋めて出力してください。

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
	fmt.Printf("%3d\n", n)
}
