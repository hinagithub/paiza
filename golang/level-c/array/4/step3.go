// 問題
// 1 行目に整数 A, B, N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数の左から A 番目の値と B 番目の値を入れ替えて、改行区切りで出力してください。
// なお、左端を 1 番目とします。

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
	a, _ := strconv.Atoi(inputs[0])
	b, _ := strconv.Atoi(inputs[1])

	scanner.Scan()
	numberStrs := strings.Fields(scanner.Text())
	numbers := make([]int, len(numberStrs))
	for i, str := range numberStrs {
		numbers[i], _ = strconv.Atoi(str)
	}

	numbers[a-1], numbers[b-1] = numbers[b-1], numbers[a-1]

	for _, num := range numbers {
		fmt.Println(num)
	}
}
