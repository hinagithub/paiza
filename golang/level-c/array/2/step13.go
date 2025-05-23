// 問題
// 1 行目に整数 N, M, K, L が与えられます。
// 2 行目以降に N 行 M 列の二次元配列が与えられます。
// 配列の K 行目 L 列目の要素を出力してください。
// 上から i 番目、左から j 番目の整数は a_ij です。

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
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	inputs := strings.Fields(scanner.Text())
	x := atoi(inputs[2])
	y := atoi(inputs[3])

	matrix := [][]int{}
	for scanner.Scan() {
		numbers := []int{}
		inputs := strings.Fields(scanner.Text())
		for _, input := range inputs {
			num := atoi(input)
			numbers = append(numbers, num)
		}
		matrix = append(matrix, numbers)
	}
	fmt.Println(matrix[x-1][y-1])

}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return num
}
