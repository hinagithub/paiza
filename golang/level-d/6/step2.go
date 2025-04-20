// 問題
// 1 ~ 100 の整数に対して、
// 3 と 5 の両方で割り切れるなら FizzBuzz を、
// 3 でのみ割り切れるなら Fizz 、
// 5 でのみ割り切れるなら Buzz を改行区切りで出力してください。
// また、どちらでも割り切れない場合は、その数字を改行区切りで出力してください。

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		remainderDividedBy3 := i % 3
		remainderDividedBy5 := i % 5

		if remainderDividedBy3 == 0 && remainderDividedBy5 == 0 {
			fmt.Println("FizzBuzz")
		} else if remainderDividedBy3 == 0 {
			fmt.Println("Fizz")

		} else if remainderDividedBy5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
