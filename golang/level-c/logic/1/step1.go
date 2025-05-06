// 問題
// 0 または 1 の整数 A と B が与えられます。 A AND B の結果を出力してください。

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
	a := atoi(inputs[0])
	b := atoi(inputs[1])

	fmt.Println(a & b)
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return num
}
