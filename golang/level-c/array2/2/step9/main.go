// 問題
// 配列 A の要素数 N と新たに作成する配列のサイズ n , 配列 A の各要素 A_1 ... A_N が与えられるので、
// 配列 A の先頭から n 要素を順に保持する配列を作成してください。
// 新たに作成する配列の要素数が A の要素数よりも大きい時は、サイズが合うように 0 を A の要素の末尾に追加してください。

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
	oldArrLength := atoi(inputs[0])
	newArrLength := atoi(inputs[1])

	// A の全要素を読み取り
	source := make([]int, oldArrLength)
	for i := 0; i < oldArrLength; i++ {
		source[i] = atoi(readline(scanner))
	}

	// 新しい配列を作成
	results := make([]int, newArrLength)
	for i := 0; i < newArrLength; i++ {
		if i < oldArrLength {
			results[i] = source[i]
		}
	}

	for _, num := range results {
		fmt.Println(num)
	}
}

func readline(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintln(os.Stdout, "failed to scan input")
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
