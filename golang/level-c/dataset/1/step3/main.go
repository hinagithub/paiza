// 問題
// 自然数 A, B, C が与えられます。(A, B, C の最大値) - (A, B, C の最小値)を答えてください。

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
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	numbers := atoiall(inputs)
	fmt.Println(max(numbers) - min(numbers))
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

func max(numbers []int) int {
	max := numbers[0]
	for _, number := range numbers[1:] {
		if number > max {
			max = number
		}
	}
	return max
}

func min(numbers []int) int {
	min := numbers[0]
	for _, number := range numbers[1:] {
		if number < min {
			min = number
		}
	}
	return min
}
