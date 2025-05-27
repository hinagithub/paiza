// 問題
// 配列 A の要素数 N と配列 A の各要素 A_1, A_2, ..., A_N が与えられるので、
// 同じ値の要素が 2 つ以上含まれないように A を加工した新たな配列 B を作成してください。
// ただし、 A に同じ値の要素が 2 つ以上あった場合、それらのうち先頭の要素を B に採用するものとします。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	length := atoi(readline(scanner))

	seen := make(map[int]bool, length)
	uniqueNums := []int{}
	for i := 0; i < length; i++ {
		num := atoi(readline(scanner))
		if !seen[num] {
			uniqueNums = append(uniqueNums, num)
			seen[num] = true
		}
	}

	for _, u := range uniqueNums {
		fmt.Println(u)
	}

}

func readline(scanner *bufio.Scanner) string {
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
