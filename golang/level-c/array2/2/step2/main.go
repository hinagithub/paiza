// 問題
// 整数型の変数を 2 つ宣言して整数 X, Y を格納したのち、
// 2 つの変数の中身を入れ替えた後の 2 つの変数の値を出力してください。

package main

import "fmt"

func main() {

	// 入力値をxとyに格納
	var x, y int
	fmt.Scan(&x, &y)

	// 2 つの変数の中身を入れ替えて出力
	y, x = x, y
	fmt.Println(x, y)
}
