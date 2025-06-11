// 問題
// 配列 A の要素数 N とその要素 A_i (1 ≦ i ≦ N) が与えられるので、A についてのかけ算表 B を出力してください。
// かけ算表は N * N の二次元配列の形式とし、B の i 行 j 列の要素 B_ij について、
// B_ij = Ai * Aj (1 ≦ i , j ≦ N) が成り立つものとします。
// A = [1,2,3] のとき B は
// 1 2 3
// 2 4 6
// 3 6 9

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
	inputs := strings.Fields(readLine(scanner))

	for i := 0; i < length; i++ {
		row := make([]string, length)
		for j := 0; j < length; j++ {
			Ai := atoi(inputs[i])
			Aj := atoi(inputs[j])
			row[j] = strconv.Itoa(Ai * Aj)
		}
		fmt.Println(strings.Join(row, " "))
	}

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
