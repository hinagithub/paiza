// 問題
// N 個の要素からなる数列 A が与えられます。
// 数列 A は昇順にソートされています。
// A の重複した要素を取り除いて昇順に出力してください。

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
	readLine(scanner) // 1行目の要素数は使用しないのでスキップ
	input := readLine(scanner)
	texts := strings.Fields(input)
	inputNumbers := atoiAll(texts)

	resultNumbers := []int{}
	for _, inputNumber := range inputNumbers {
		if !found(resultNumbers, inputNumber) {
			resultNumbers = append(resultNumbers, inputNumber)
		}
	}

	resultTexts := itoaAll(resultNumbers)
	fmt.Println(strings.Join(resultTexts, " "))

}

func readLine(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scanner err")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoiAll(texts []string) []int {
	numbers := []int{}
	for _, text := range texts {
		number, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(os.Stderr, "atoi err", err)
			os.Exit(1)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func itoaAll(numbers []int) []string {
	texts := []string{}
	for _, number := range numbers {
		text := strconv.Itoa(number)
		texts = append(texts, text)
	}
	return texts
}

func found(numbers []int, target int) bool {
	for _, number := range numbers {
		if number == target {
			return true
		}
	}
	return false
}
