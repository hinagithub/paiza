// 生徒の身長が A_1, ...., A_N であるような N 人のクラスで二人三脚の代表を決めることにしました。
// 代表には、身長の差が最も小さいような 2 人を選出することにしました。
// 選ばれた 2 人の身長を昇順に出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	personCount := atoi(readLine(scanner))

	heights := make([]int, personCount)
	for i := 0; i < personCount; i++ {
		heights[i] = atoi(readLine(scanner))
	}

	sort.Ints(heights)

	minDiff := abs(heights[1] - heights[0])
	closestHeight1 := heights[0]
	closestHeight2 := heights[1]

	for i := 1; i < personCount-1; i++ {
		currentDiff := abs(heights[i+1] - heights[i])
		if currentDiff < minDiff {
			minDiff = currentDiff
			closestHeight1 = heights[i]
			closestHeight2 = heights[i+1]
		}
	}

	fmt.Println(closestHeight1)
	fmt.Println(closestHeight2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		fmt.Fprintf(os.Stderr, "failed to convert to int: %v\n", err)
		os.Exit(1)
	}
	return num
}
