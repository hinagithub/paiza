// 問題
// 文字列 S と整数 i , j が与えられるので、 S の i 文字目から j 文字目までの部分文字列を出力してください。

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
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	text := scanner.Text()
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "scan err")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())

	start, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi err")
		os.Exit(1)
	}
	end, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi err")
		os.Exit(1)
	}
	fmt.Println(text[start-1 : end])
}
