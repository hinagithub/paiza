// 問題
// 1 行目に整数 N と整数 M が与えられます。
// 2 行目以降に N 行 M 列の配列が与えられます。上から i 番目、左から j 番目の整数は a_ij です。
// 全要素を各行ずつ半角スペース区切りで出力し、行の終わりで改行してください。

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 1行目は読み飛ばす
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
