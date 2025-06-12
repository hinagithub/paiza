// 問題
// 整数 N が与えられるので、2 以上 N 以下の素数の個数を求めてください。
// 素数とはの約数が 1 と X のみであるような整数 X のことを指します。

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}

	}
	fmt.Println(count)
}

func isPrime(num int) bool {
	prime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			prime = false
			break
		}
	}
	return prime
}
