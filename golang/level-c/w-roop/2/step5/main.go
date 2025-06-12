// 問題
// 普通の鳩時計は 1 時間に 1 回しか鳴かないのでつまらないと思ったあなたは、
// 鳩時計を改造してスーパー鳩時計を作りました。
// このスーパー鳩時計は時刻が x 時 y 分のとき x + y が 3の倍数のとき"FIZZ"、
// 5 の倍数のとき"BUZZ", 3の倍数かつ5の倍数のとき "FIZZBUZZ" と鳴き、
// これらのいずれにも当てはまらなかった場合は鳴きません。
// なお、0 は 3 の倍数かつ 5 の倍数であるとします。
// 0 時 0 分　〜 23 時 59 分 の各分のスーパー鳩時計の様子を出力してください。

package main

import "fmt"

func main() {

	// 24時間
	for i := 0; i < 24; i++ {

		// 60分
		for j := 0; j < 60; j++ {
			fizzbuzz(i, j)
		}
	}
}

func fizzbuzz(x int, y int) {
	z := x + y
	if z%5 == 0 && z%3 == 0 {
		fmt.Println("FIZZBUZZ")
	} else if z%3 == 0 {
		fmt.Println("FIZZ")
	} else if z%5 == 0 {
		fmt.Println("BUZZ")
	} else {
		// 何もない時は改行のみ出力
		fmt.Println()
	}
}
