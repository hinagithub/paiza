// 問題
// 数値を表す文字列 S が与えられるので、 S - 813 の値を求めてください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const minus = 813

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "scan err")
		os.Exit(1)
	}
	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err", err)
		os.Exit(1)
	}

	fmt.Println(number - 813)

}
