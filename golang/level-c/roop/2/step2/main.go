// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) と、整数 A, B が与えられます。
// この数列の a_A から a_B までの和を計算し、出力してください。

// N A B
// a_1 a_2 ... a_N
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
	inputs := strings.Fields(readline(scanner))
	// n := atoi(inputs[0]) // 要素数は使用しない
	a := atoi(inputs[1])
	b := atoi(inputs[2])

	strNumbers := strings.Fields(readline(scanner))

	sum := 0
	for i := a - 1; i < b; i++ {
		num := atoi(strNumbers[i])
		sum += num
	}

	fmt.Println(sum)

}

func readline(scanner *bufio.Scanner) string {
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
