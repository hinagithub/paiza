// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ..., A_N ,
// 交換する A の要素番号　X, Y が与えられるので、A_X と A_Y を入れ替えた後の A を出力してください。

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

	// 要素数を取得
	length, err := strconv.Atoi(readline(scanner))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}

	// 数値配列を取得
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = atoi(readline(scanner))
	}
	inputs := strings.Fields(readline(scanner))

	// xとyの要素番号を取得
	indexX := atoi(inputs[0]) - 1
	indexY := atoi(inputs[1]) - 1

	// xとyを入れ替えて出力
	numbers[indexX], numbers[indexY] = numbers[indexY], numbers[indexX]
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "failed to scan input")
		os.Exit(1)
	}
	return scanner.Text()
}
