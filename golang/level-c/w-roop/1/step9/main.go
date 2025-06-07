// 整数 N が与えられるので、次の処理を N 回してください。
// 配列のサイズ K とその要素 A1 ... AK が与えられるので、全ての要素の和を求めて出力してください。

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

	// 行ごとの和を出力 (要素0は要素数なので使用しない)
	for i := 0; i < length; i++ {
		row := strings.Fields(readLine(scanner))
		sam := 0
		
		for _, text := range row[1:] {
			num := atoi(text)
			sam += num
		}
		fmt.Println(sam)
	}
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
