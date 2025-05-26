// 問題
// 整数 A, B, C, D が与えられます。
// ((A+B)*C)^2 mod D を計算した結果を出力してください。
// ここで、X mod D とは X を D で割った余りを指します。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())

	a := mustAtoi(inputs[0])
	b := mustAtoi(inputs[1])
	c := mustAtoi(inputs[2])
	d := mustAtoi(inputs[3])

	sumProduct := (a + b) * c
	squared := sumProduct * sumProduct
	result := squared % d

	fmt.Println(result)
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid integer: %s\n", s)
		os.Exit(1)
	}
	return n
}
