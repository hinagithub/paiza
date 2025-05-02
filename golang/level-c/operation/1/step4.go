// 問題
// 整数 A, B, C が与えられます。
// A * Aの計算結果 X とB * B + C * Cの計算結果 Y を半角スペース区切りで出力してください。

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
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scan err")
		os.Exit(1)
	}

	inputs := strings.Fields(scanner.Text())
	a := aToI(inputs[0])
	b := aToI(inputs[1])
	c := aToI(inputs[2])

	x := a * a
	y := b*b + c*c
	fmt.Println(x, y)

}

func aToI(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return number
}
