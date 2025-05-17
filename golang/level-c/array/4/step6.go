// 問題
// 1 行目に整数 N が与えられます。
// 2 行目に N 個の整数 a_1, a_2, ..., a_N が与えられます。
// N 個の整数を小さい順にソートし、改行区切りで出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行目は使わないので読み飛ばす
	scanner.Scan()
	inputs := strings.Fields(scanner.Text())
	numbers := atoiAll(inputs)
	// sortedNumbers := sort(numbers)

	// for _, num := range sortedNumbers {
	// 	fmt.Println(num)
	// }

	// 標準ライブラリ使ってできた
	sort.Ints(numbers)
	for _, num := range numbers {
		fmt.Println(num)
	}

}

// func sort(numbers []int) []int {

// 	length := len(numbers)

// 	for i := 0; i < len(numbers)-1; i++ {
// 		maxIndex, maxNum := findMaxIndex(numbers[:length])
// 		numbers[maxIndex] = numbers[length-1]
// 		numbers[length-1] = maxNum
// 		length--
// 	}
// 	return numbers

// }

// func findMaxIndex(numbers []int) (int, int) {
// 	if len(numbers) == 0 {
// 		return -1, 0
// 	}

// 	maxIndex := 0
// 	maxValue := numbers[0]

// 	for i, num := range numbers {
// 		if num > maxValue {
// 			maxValue = num
// 			maxIndex = i
// 		}
// 	}

// 	return maxIndex, maxValue
// }

func atoiAll(strNumbers []string) []int {
	numbers := make([]int, len(strNumbers))
	for i, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "atoi err: %d", err)
			os.Exit(1)
		}
		numbers[i] = num
	}
	return numbers
}
