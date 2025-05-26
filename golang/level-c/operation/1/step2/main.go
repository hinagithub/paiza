// 問題
// 整数 A, B が与えられます。
// A と B の差 D と積 P を半角スペース区切りで出力してください。

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
	a := atoi(inputs[0])
	b := atoi(inputs[1])
	fmt.Println(a-b, a*b)
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return num
}
