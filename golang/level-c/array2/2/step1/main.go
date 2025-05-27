// 問題
// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ...,
// A_N が与えられるので、A の要素の順序を逆にした配列 B を作成し、出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	length := atoi(readline(scanner))

	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[length-1-i] = atoi(readline(scanner))
	}
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "faild to scan inputs")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "faild to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
