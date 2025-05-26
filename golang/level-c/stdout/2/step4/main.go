// 1 から 10 までの数値をすべて、半角スペース区切りで出力してください。
// ただし、末尾に半角スペースを出力してはいけません。

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i)
		if i != 10 {
			fmt.Print(" ")
		}
	}
}
