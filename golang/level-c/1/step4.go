// 以下の 10 個の整数を改行区切りで出力してください。

package main

import "fmt"

func main() {
	numbers := [10]int{813, 1, 2, 923874, 23648, 782356, 3256, 2342, 24324, 112}
	for _, number := range numbers {
		fmt.Println(number)
	}
}
