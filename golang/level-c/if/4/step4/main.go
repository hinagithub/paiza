// 問題
// 長さ N の数列Aが与えられます。 1 つ目の要素から最も左にある奇数の要素の手前までの値の和を求めてください。
// 形式的には、A_iが奇数かつ、区間 [A_1, A_{i - 1}] がすべて偶数であるとき、\sum_{j=1}^{i-1}{A_j}を求めてください。

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

	sum := 0
	for i := 0; i < length; i++ {
		num := atoi(strNumbers[i])

		// 奇数を見つけたら終了
		if num%2 != 0 {
			break
		}
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
