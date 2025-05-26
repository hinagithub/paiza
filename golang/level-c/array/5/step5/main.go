// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、N 個の整数の平均以上の数をすべて、入力された順に改行区切りで出力してください。

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
	length := atoi(scanner.Text())

	scanner.Scan()

	inputs := strings.Fields(scanner.Text())

	sum := 0
	numbers := []int{}
	for _, input := range inputs {
		num := atoi(input)
		sum = sum + num
		numbers = append(numbers, num)
	}

	// 平均値を割り出す
	average := float64(sum) / float64(length)

	for _, num := range numbers {
		if float64(num) >= average {
			fmt.Println(num)
		}
	}
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return num
}
