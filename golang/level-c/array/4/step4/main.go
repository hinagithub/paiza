// 問題
// 1 行目に整数 A, B, N （A ≦ B）が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、左から A 番目から B 番目までの数を抜き出し、改行区切りで出力してください。
// なお、左端を 1 番目とします。

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
	inputs := strings.Fields(scan(scanner))
	start, _ := strconv.Atoi(inputs[0])
	end, _ := strconv.Atoi(inputs[1])

	for _, numStr := range strings.Fields(scan(scanner))[start-1 : end] {
		fmt.Println(numStr)
	}

}

func scan(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan error")
		os.Exit(1)
	}
	return scanner.Text()
}
