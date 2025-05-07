// 問題
// 数値 X , Y が与えられるので、X + Y の計算結果の先頭から N 文字目の数字を出力してください。

// 例
// ・ X = 813 , Y = 813 , N = 1 のとき
// X + Y = 1626 の 1 文字目である 1 を出力してください。

// ・ X = -813 , Y = 813 , N = 1 のとき
// X + Y = 0 の 1 文字目である 0 を出力してください。

// ・ X = 10000 , Y = -813 , N = 3 のとき
// X + Y = 9187 の 3 文字目である 8 を出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	X := atoi(read(scanner))
	Y := atoi(read(scanner))
	N := atoi(read(scanner))

	result := X + Y
	strResult := strconv.Itoa(result)
	if N < 0 || N > len(strResult) {
		fmt.Fprintf(os.Stderr, "index out of range")
		os.Exit(1)
	}
	fmt.Println(string(strResult[N-1]))

}

func read(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "scan err")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return num
}
