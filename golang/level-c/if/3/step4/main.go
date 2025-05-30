// 問題
// ある月の 1 日は日曜日、 2 日は月曜日...です。X日は何曜日でしょう。
package main

import (
	"fmt"
)

func main() {
	weekdays := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	var x int
	fmt.Print("何日？: ")
	fmt.Scan(&x)

	// 1日がSunなので、x-1 にして 0-index に合わせる
	fmt.Println(weekdays[(x-1)%7])
}
