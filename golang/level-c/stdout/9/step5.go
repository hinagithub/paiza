// 問題
// 自然数 N, M と N 個の自然数 A_1, A_2, ..., A_N が与えられます。
// それぞれの数値が M けたになるよう数値の前に半角スペースを埋めて、改行区切りで出力してください。

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
	len, _ := strconv.Atoi(inputs[0])
	width, _ := strconv.Atoi(inputs[1])
	for i := 0; i < len; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		format := fmt.Sprintf("%%%dd\n", width)
		fmt.Printf(format, n)
	}
}
