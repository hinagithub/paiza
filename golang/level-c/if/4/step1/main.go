// 問題
// 整数Nが与えられます。Nのけた数を出力してください。

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(len(strconv.Itoa(n)))
}

// 別のやり方

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	if n == 0 {
// 		fmt.Println(1)
// 		return
// 	}

// 	count := 0
// 	for n != 0 {
// 		n /= 10
// 		count++
// 	}
// 	fmt.Println(count)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	if n == 0 {
// 		fmt.Println(1)
// 		return
// 	}

// 	digits := int(math.Log10(float64(n))) + 1
// 	fmt.Println(digits)
// }
