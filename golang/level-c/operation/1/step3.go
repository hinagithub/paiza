// 問題
// 整数 A = 437,326、 B = 9,085 とします。以下の二つの演算の結果を半角スペース区切りで出力してください。

// 1. A を B で割った商
// 2. A を B で割った余り

package main

import "fmt"

func main() {
	const (
		a = 437326
		b = 9085
	)
	fmt.Println(a/b, a%b)
}
