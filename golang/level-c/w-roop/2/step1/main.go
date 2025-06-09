// 問題
// N 行 K 列の行列 A の i 行 j 列 の要素 A_ij を A_ji とした K 行 N 列の行列を元の配列 A の転置行列と言います。
// 行列 A についての情報が与えられるので、A の転置行列を出力してください。
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

	// N（行数）, K（列数）を取得
	inputs := strings.Fields(readLine(scanner))
	rowLen := atoi(inputs[0])
	colLen := atoi(inputs[1])
	matrix := make([][]int, rowLen)
	for i := 0; i < rowLen; i++ {
		row := strings.Fields(readLine(scanner))
		matrix[i] = make([]int, colLen)
		for j := 0; j < colLen; j++ {
			matrix[i][j] = atoi(row[j])
		}
	}

	// 転置行列を作成
	transposed := make([][]int, colLen)
	for i := 0; i < colLen; i++ {
		transposed[i] = make([]int, rowLen)
		for j := 0; j < rowLen; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}

	for _, row := range transposed {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
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
