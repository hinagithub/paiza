// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) が与えられます。
// この数列の要素を逆順に、改行区切りで出力してください。

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
	length := atoi(readline(scanner))
	strNumbers := strings.Fields(readline(scanner))

	for i := 0; i < length; i++ {
		num := strNumbers[len(strNumbers)-1-i]
		fmt.Println(num)
	}
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
