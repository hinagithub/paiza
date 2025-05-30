// 問題
// ある月の 1 日は日曜日、 2 日は月曜日...です。X日は何曜日でしょう。

package main

import (
	"fmt"
)

func main() {

	const (
		Sun = "Sun"
		Mon = "Mon"
		Tue = "Tue"
		Wed = "Wed"
		Thu = "Thu"
		Fri = "Fri"
		Sat = "Sat"
	)
	weekdays := map[int]string{
		1:  Sun,
		2:  Mon,
		3:  Tue,
		4:  Wed,
		5:  Thu,
		6:  Fri,
		7:  Sat,
		8:  Sun,
		9:  Mon,
		10: Tue,
		11: Wed,
		12: Thu,
		13: Fri,
		14: Sat,
		15: Sun,
		16: Mon,
		17: Tue,
		18: Wed,
		19: Thu,
		20: Fri,
		21: Sat,
		22: Sun,
		23: Mon,
		24: Tue,
		25: Wed,
		26: Thu,
		27: Fri,
		28: Sat,
		29: Sun,
		30: Mon,
	}

	var x int
	fmt.Scan(&x)

	fmt.Println(weekdays[x])

}
