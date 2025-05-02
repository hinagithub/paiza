// 問題
// 以下のアルゴリズムを実行してください。
// 変数 N を 10,000 で初期化する
// N を 361 で割った商を N へ代入する
// N を 28 で割った余りを N へ代入する
// N を出力する

package main

import "fmt"

func main() {

	n := 10000
	n = n / 361
	n = n % 28
	fmt.Println(n)
}
