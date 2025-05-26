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
	scanner.Scan() // 要素数(n)は不要なのでスキップ
	const count = 10

	// 要素を取得
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	numbers := atoiall(inputs)

	resultNumbers := make([]int, count)
	for i := 0; i < count; i++ {
		resultNumbers = append(resultNumbers, countOccurrences(numbers, i))
	}

	results := itoaall(resultNumbers)
	fmt.Println(strings.Join(results, " "))

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

func itoaall(numbers []int) []string {
	texts := []string{}
	for _, number := range numbers {
		text := strconv.Itoa(number)
		texts = append(texts, text)
	}
	return texts
}

func countOccurrences(numbers []int, target int) int {
	count := 0
	for _, number := range numbers {
		if number == target {
			count++
		}
	}
	return count
}
