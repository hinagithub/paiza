// 問題
// paiza の入社試験では 科目 1 〜 5 の 5 科目のテストが課せられており、
// それぞれの科目には重みが設定されています。受験者の得点は各科目の (とった点数) * (科目の重み) となります。
// 5 科目の得点の合計が最も高かった受験者の得点を求めてください。

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
	weightStrs := strings.Fields(readline(scanner))
	weights := atoiAll(weightStrs)

	max := 0
	for i := 0; i < length; i++ {
		scores := strings.Fields(readline(scanner))
		totalScore := 0
		for i, score := range scores {
			totalScore = totalScore + atoi(score)*weights[i]
		}
		if totalScore > max {
			max = totalScore
		}
	}

	fmt.Println(max)
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}

func atoiAll(texts []string) []int {
	numbers := make([]int, len(texts))
	for i, text := range texts {
		numbers[i] = atoi(text)
	}
	return numbers
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}
