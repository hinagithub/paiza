// 問題
// 自然数 N, M が与えられます。
// N が M けたになるよう数値の前に半角スペースを埋めて出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(inputs[0])
	width, _ := strconv.Atoi(inputs[1])
	format := fmt.Sprintf("%%%dd\n", width)
	fmt.Printf(format, n)
}
