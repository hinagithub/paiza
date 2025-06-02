// 問題
// 長さ N の数列 a (a_1, a_2, ..., a_N) が与えられます。
// この数列の何番目に 1 があるか出力してください。
// 数列の 1 つ目の要素を 1 番目とし、数列には必ず 1 がひとつだけ含まれることとします。

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
	readline(scanner) // 1行目の要素数nは使用しないのでスキップ
	strNumbers := strings.Fields(readline(scanner))
	for i, strNum := range strNumbers {
		num := atoi(strNum)
		if num == 1 {
			fmt.Println(i + 1)
			return
		}
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
