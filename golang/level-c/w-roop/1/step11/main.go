// 配列 A と B についての情報が与えられるので、(A の 1 つの要素) × (B の 1 つの要素) の最大値を求めてください。

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
	_ = readLine(scanner)
	arr1 := strings.Fields(readLine(scanner))
	arr2 := strings.Fields(readLine(scanner))

	max := -1 << 31 // intの最小値で初期化
	for _, aStr := range arr1 {
		a := atoi(aStr)
		for _, bStr := range arr2 {
			b := atoi(bStr)
			product := a * b
			if product > max {
				max = product
			}
		}
	}
	fmt.Println(max)
}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
