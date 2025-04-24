// 問題
// 自然数 N が与えられます。
// N 個の自然数が与えられるので、それぞれを数値 M_i について以下の処理を行ってください。
// * M_i が 3 けたになるよう数値の前に半角スペースを埋めて出力してください。

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
	len, _ := strconv.Atoi(sc.Text())
	for i := 0; i < len; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		fmt.Printf("%3d\n", n)
	}
}
