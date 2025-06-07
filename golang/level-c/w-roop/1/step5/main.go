// 整数 N , K と N 行 K 列 の二次元配列 A が与えられるので、その配列をそのまま出力してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行目の要素数は使用しない
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
