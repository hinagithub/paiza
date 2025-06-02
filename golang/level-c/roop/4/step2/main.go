// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) が与えられます。
// この数列の最小値を出力してください。

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
	length := atoi(readLine(scanner))
	strNumbers := strings.Fields(readLine(scanner))
	min := atoi(strNumbers[0])
	for i := 0; i < length; i++ {

		num := atoi(strNumbers[i])
		if num < min {
			min = num
		}
	}
	fmt.Println(min)
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
