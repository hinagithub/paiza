// 問題
// 整数Nが与えられます。
// Nが 3 で割り切れる場合はFizz、Nが 5 で割り切れる場合はBuzz、
// Nが 3 と 5 の両方で割り切れる場合はFizzやBuzzの代わりにFizzBuzzを出力してください。
// ただし、Nが 3 の倍数でも 5 の倍数でもないときはNをそのまま出力してください。

package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(x)
	}
}
