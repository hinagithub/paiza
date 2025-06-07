// 整数 N , K と N 行 K 列 の二次元配列 A が与えられます。 A の要素のうち、最大の要素の値を出力してください。

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
	inputs := strings.Fields(readLine(scanner))
	length := atoi(inputs[0])

	max := 0
	// 2重配列を操作し1を見つけたら行例を出力
	for i := 0; i < length; i++ {
		row := strings.Fields(readLine(scanner))
		for _, text := range row {
			num := atoi(text)
			if num > max {
				max = num
			}
		}
	}

	fmt.Println(max)

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
		fmt.Fprintf(os.Stderr, "failed to atoi: %v \n", err)
		os.Exit(1)
	}
	return num
}
