// 問題
// 1 行目に整数 K と 整数 L が与えられます。
// 2 行目以降に 3 行 3 列の二次元配列が与えられます。
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
	inputs := strings.Fields(readline(scanner))
	x := atoi(inputs[0])
	y := atoi(inputs[1])

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

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println("scan err")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return num
}
