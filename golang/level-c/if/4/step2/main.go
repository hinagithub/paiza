// 整数N, A, B ( - 99 ≦ N, A, B ≦ 100 ) があります。
// 以下の 2 つの操作をそれぞれ 1 回ずつおこなったとき、Nを 0 にできる場合はYESを、できない場合はNOを出力してください。
// 1. NにAを足す、またはNからAを引く
// 2. NにBを足す、またはNからBを引く

package main

import "fmt"

func main() {

	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if addadd(n, a, b) == 0 || addsub(n, a, b) == 0 || subadd(n, a, b) == 0 || subsub(n, a, b) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func addadd(n int, a int, b int) int {
	return n + a + b
}
func addsub(n int, a int, b int) int {
	return n + a - b
}
func subadd(n int, a int, b int) int {
	return n - a + b
}
func subsub(n int, a int, b int) int {
	return n - a - b
}
