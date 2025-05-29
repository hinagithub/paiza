// 問題
// 長さ N の数列Aが与えられます。
// Aの中に 0 が含まれていない場合はYESを、 0 が含まれている場合はNOを出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行目の要素数は使わないのでスキップ

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to atoi : %v \n", err)
			os.Exit(1)
		}

		if n == 0 {
			fmt.Println("NO")
			os.Exit(0)
		}
	}
	fmt.Println("YES")
}
