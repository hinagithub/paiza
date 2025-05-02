// 問題
// 整数 A, B が与えられます。以下のアルゴリズムを実行してください。
// 変数 N を 10,000 で初期化する
// N を A で割った商を N へ代入する
// N を B で割った余りを N へ代入する
// N を出力する

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
	inputs := strings.Fields(read(scanner))
	a := atoi(inputs[0])
	b := atoi(inputs[1])

	n := 10000
	n = n / a
	n = n % b
	fmt.Println(n)

}

func read(scanner *bufio.Scanner) string {
	if !scanner.Scan() {
		fmt.Println(os.Stderr, "scan err")
		os.Exit(1)
	}
	return scanner.Text()
}

func atoi(text string) int {
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(os.Stderr, "atoi err", err)
		os.Exit(1)
	}
	return num
}
