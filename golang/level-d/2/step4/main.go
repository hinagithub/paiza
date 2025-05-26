// 問題
// 変数 N を 0 で初期化する
// N に 3,286 を足した値を N へ代入する
// N に 4,736 をかけた値を N へ代入する
// N を 12,312 で割った余りを N へ代入する
// N を出力する

package main

import "fmt"

func main() {
	n := 0
	n = n + 3286
	n = n * 4736
	n = n % 12312
	fmt.Println(n)
}
