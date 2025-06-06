// 問題
// N 個の整数 a_1, a_2, ..., a_N が与えられます。
// a_1, a_2, ..., a_N のうち、1 がある位置を先頭から順に改行区切りで出力してください。
// a_1 を 1 番目とし、a_1, a_2, ..., a_N には少なくとも 1 個は 1 が含まれます。

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
	_ = readLine(scanner) // 1 行目の要素数は使用しない
	inputs := strings.Fields(readLine(scanner))

	// 配列を操作し1の時だけ要素番号(i + 1)を出力
	for i, input := range inputs {
		num := atoi(input)
		if num == 1 {
			fmt.Println(i + 1)
		}
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

func atoiAll(texts []string) []int {
	nums := []int{}
	for _, text := range texts {
		num := atoi(text)
		nums = append(nums, num)
	}
	return nums
}
