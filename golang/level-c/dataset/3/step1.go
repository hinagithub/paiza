// 問題
// 0 以上 9 以下の整数が N 個与えられます。
// 各数値の出現回数を求め、
// 「0」の出現回数、「1」の出現回数、...「9」の出現回数、
// をこの順に半角スペース区切りで1行に出力してください。

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
	scanner.Scan() // 要素数(N)は使用しない
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	numbers := atoiall(inputs)

	results := []int{}
	for i, number := range numbers {

	}
}

func atoiall(texts []string) []int {
	numbers := []int{}
	for _, text := range texts {
		number, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("atoi err")
			os.Exit(1)
		}
		numbers = append(numbers, number)

	}
	return numbers
}
