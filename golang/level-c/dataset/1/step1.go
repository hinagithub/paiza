// 問題
// 要素数 N の数列 A と数値 M が与えられます。A の M 番目の値を出力してください。

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
	targetIndex, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Println("target index not found...")
		os.Exit(1)
	}
	targetIndex--

	scanner.Scan()
	texts := strings.Fields(scanner.Text())
	fmt.Println(texts[targetIndex])

}
