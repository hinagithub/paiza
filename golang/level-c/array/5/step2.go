// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数のうち、最大の数と最小の数を半角スペース区切りで出力してください。
// N 個の整数を大きい順や小さい順に並び替える操作を考えて解いてみましょう。

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
	scanner.Scan() // 1行目は使わないので読み飛ばす
	scanner.Scan()
	strNumbers := strings.Fields(scanner.Text())
	numbers := atoiAll(strNumbers)
	fmt.Println(findMax(numbers), findMin(numbers))

}

func atoiAll(strs []string) []int {
	numbers := []int{}
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "atoi err: %d", err)
			os.Exit(1)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func findMax(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func findMin(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}
