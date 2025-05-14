// 問題
// 1 行目に整数 N, M が与えられます。
// 2 行目に M 個の整数 a_1, a_2, ..., a_M が与えられます。
// 整数 N が、M 個の整数の左から何番目にあるか出力してください。
// 左端を 1 番目とし、N は M 個の整数に必ず 1 つだけ含まれるものとします。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scan(scanner)
	inputs := strings.Fields(scanner.Text())
	target := inputs[0]
	scan(scanner)
	for i, str := range strings.Fields(scanner.Text()) {
		if str == target {
			fmt.Println(i + 1)
		}
	}
}

func scan(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
}
