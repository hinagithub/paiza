// 問題
// N 個の整数 a_1, a_2, ..., a_N が与えられます。
// a_i に i を足したとき、N 個の整数の最小値を出力してください。

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
	_ = readLine(scanner) // 1行目の要素数は使用しない
	inputs := strings.Fields(readLine(scanner))
	numbers := atoiAll(inputs)

	// 数値配列を操作し i を足した数が最大の値を出力する
	min := numbers[0] + 1
	for i, number := range numbers {
		adjusted := number + i + 1
		if adjusted < min {
			min = adjusted
		}
	}
	fmt.Println(min)

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

func atoiAll(texts []string) []int {
	nums := []int{}
	for _, text := range texts {
		num := atoi(text)
		nums = append(nums, num)
	}
	return nums
}
