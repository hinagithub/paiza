// 問題
// 配列 A の要素数 N と整数 K , 配列 A の各要素 A_1 ... A_N が与えられるので、
// K 以上であるような配列 A の要素のみを要素として持つ配列 B を作成し、その要素を出力してください。
//  B の要素の順は A と同じにしてください。

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
	length := atoi(inputs[0])
	border := atoi(inputs[1])

	numbers := []int{}
	for i := 0; i < length; i++ {
		num := atoi(readline(scanner))
		if num >= border {
			numbers = append(numbers, num)
		}
	}

	for _, num := range numbers {
		fmt.Println(num)
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
