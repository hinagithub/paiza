// 問題
// 長さ N の数列Aが与えられます。
// この数列に含まれる偶数の要素の個数と、奇数の要素の個数を答えてください。

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

	evenCount := 0
	oddCount := 0
	for i := 0; i < length; i++ {
		num := atoi(strNumbers[i])
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Println(evenCount, oddCount)
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
		fmt.Fprintf(os.Stderr, "failed to atoi: %v\n", err)
		os.Exit(1)
	}
	return num
}
